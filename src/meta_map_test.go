package main

import (
	"errors"
	"github.com/golang-collections/collections/stack"
)

/*func iterate(collection []reflect.Value){}

func assign(value reflect.Value, variable reflect.Value){}

func dotAccess(path string, object reflect.Value){}

func callFunction(args []reflect.Value){}*/

type decision func(blockStack *stack.Stack, match Section) (error, Section)

func isOpen (blockStack *stack.Stack, match Section) (error, Section) {
	blockStack.Push(match)
	return nil, Section{}
}

func isClosed (blockStack *stack.Stack, match Section) (error, Section) {
	stackLength := blockStack.Len()

	if stackLength > 1 {
		blockStack.Pop()
	} else if stackLength == 1 {
		pop := blockStack.Pop()

		startMatch, ok := pop.(Section)
		if ok {
			return nil, Section{startMatch.start, match.end, startMatch.source}
		}

	}

	return errors.New("found a closing brace with no open brace"), Section{}
}



func something(match Section, blockStack *stack.Stack, blockPattern *OpenClosedPattern) (error, Section) {
	symbol := match.toString()

	decisions := map[string]decision{ "{" : isOpen, "}" : isClosed }

	return decisions[symbol](blockStack, match)
}

func altFindFirstCodeBlock(originalSection *Section, blockPattern *OpenClosedPattern) (Section, error) {

	matches := originalSection.findAll(blockPattern.whole)
	blockStack := stack.New()

	for _, match := range matches {
		err, section := something(match, blockStack, blockPattern)

		if err != nil {
			return Section{}, err
		}

		if !section.isEmpty() {
			return section, err
		}

	}

	return Section{}, errors.New("no matches found in this section")

}
