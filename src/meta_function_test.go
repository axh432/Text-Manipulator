// +build test

package main

import (
	"github.com/golang-collections/collections/stack"
)

var matches []Section
var blockStack = stack.New()
var blockPattern = newOpenClosedPatternString("{", "}")

func pushOpeningMatch(blockStack *stack.Stack, match *Section) {
	blockStack.Push(match)
}

//the stack is always going to contain sections so don't need to handle the error if the cast fails
func popBlockStack() Section {
	return blockStack.Pop().(Section)
}

func createCodeBlock(match *Section) Section {
	startMatch := popBlockStack()
	return Section{startMatch.start, match.end, startMatch.source}
}

type Action func() Section

type Branch struct {
	cond   bool
	action Action
}

func createEmptySectionAction() Action {
	return func() Section {
		return Section{}
	}
}

func createCodeBlockAction(match *Section) Action {
	return func() Section {
		return createCodeBlock(match)
	}
}

func createPushOpeningMatchAction(blockStack *stack.Stack, match *Section) Action {
	return func() Section {
		pushOpeningMatch(blockStack, match)
		return Section{}
	}
}

func createDetermineIfBlockIsFoundAction(blockStack *stack.Stack, match *Section) Action {
	return func() Section {
		return determineIfBlockIsFound(blockStack, match)
	}
}

//this needs to have a default statement! never call without one at the end.
func Switch(branches []Branch) Section {
	for _, branch := range branches {
		if branch.cond {
			return branch.action()
		}
	}
	//this should never be called if the end branch is always the default statement.
	branchesSize := len(branches)
	return branches[branchesSize-1].action()
}

//stack length is exactly 1 : create the code block
//stack length is > 1 : pop another off the stack, return an empty section
//stack length is < 1 : there is no closing brace, return an empty section
func determineIfBlockIsFound(blockStack *stack.Stack, match *Section) Section {
	stackLength := blockStack.Len()
	return Switch([]Branch{
		{stackLength > 1, popBlockStack},
		{stackLength == 1, createCodeBlockAction(match)},
		{true, createEmptySectionAction()},
	})
}

func determineBasedOnMatch(match *Section) Section {
	symbol := match.toString()
	return Switch([]Branch{
		{blockPattern.isOpen(symbol), createPushOpeningMatchAction(blockStack, match)},
		{blockPattern.isClosed(symbol), createDetermineIfBlockIsFoundAction(blockStack, match)},
		{true, createEmptySectionAction()},
	})
}

func Filter(sections []Section, function func(section *Section) bool) (prod []Section) {
	for _, section := range sections {
		if function(&section) {
			prod = append(prod, section)
		}
	}
	return
}

func Map(sections []Section, function func(match *Section) Section) (prod []Section) {
	for _, section := range sections {
		prod = append(prod, function(&section))
	}
	return
}

func loopThroughMatches(sections []Section) Section {

	isSectionNotEmpty := func(section *Section) bool {
		return !section.isEmpty()
	}

	return Filter(Map(sections, determineBasedOnMatch), isSectionNotEmpty)[0]

	/*for _, match := range matches {

		codeBlock := determineBasedOnMatch(&match)

		if !codeBlock.isEmpty() {
			return codeBlock
		}

	}

	return Section{}*/
}

func (s *Section) findFirstCodeBlock() Section {
	matches = s.findAll(blockPattern.whole)
	blockStack = stack.New()
	return loopThroughMatches(matches)
}
