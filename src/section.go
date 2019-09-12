package main

import (
	"errors"
	"fmt"
	"github.com/golang-collections/collections/stack"
	. "regexp"
)

type Section struct {
	start  int
	end    int
	source string
}

func createSectionFromString(source string) Section {
	return Section{0, len(source), source}
}

func createSectionFromMatch(match []int, offset int, source string) Section {
	return Section{offset + match[0], offset + match[1], source}
}

func (s *Section) toString() string {
	return s.source[s.start:s.end]
}

func (s *Section) find(pattern *Regexp) (Section, error) {
	start := s.start

	sectionAsString := s.toString()

	match := pattern.FindStringIndex(sectionAsString)

	if match == nil {
		return Section{}, errors.New("no matches found in this section")
	}

	return createSectionFromMatch(match, start, s.source), nil
}

func (s *Section) findAll(pattern *Regexp) (matchingSections []Section) {

	start := s.start

	sectionAsString := s.toString()

	matches := pattern.FindAllStringIndex(sectionAsString, -1)

	if matches == nil {
		return matchingSections
	}

	for _, match := range matches {
		matchingSections = append(matchingSections, createSectionFromMatch(match, start, s.source))
	}

	return matchingSections
}

func (s *Section) findAllStartEndPattern(startPattern *Regexp, endPattern *Regexp) (matchingSections []Section) {

	startMatches := s.findAll(startPattern)

	for _, startMatch := range startMatches {
		restOfSection := Section{startMatch.start, s.end, s.source}

		endMatch, err := restOfSection.find(endPattern)

		if err == nil {
			matchingSections = append(matchingSections, Section{startMatch.start, endMatch.end, startMatch.source})
		}
	}
	return
}

func (s *Section) findLinesContaining(pattern *Regexp) []Section {

	lines := []Section{}
	matches := s.findAll(pattern)

	for _, match := range matches {
		lines = append(lines, match.getLines()...)
	}

	return lines

}

func (s *Section) getFirstNewLine() Section {
	newLineRegex := MustCompile(`\n`)

	fileUpToSection := Section{ 0, s.start, s.source }

	newLinesUpToSection := fileUpToSection.findAll(newLineRegex)

	numNewLinesUpToSection := len(newLinesUpToSection)

	var firstNewLine Section

	if numNewLinesUpToSection == 0 {
		firstNewLine = Section {0, 0, s.source}
	}else{
		firstNewLine = newLinesUpToSection[numNewLinesUpToSection-1]
	}

	return firstNewLine
}

func (s *Section) getLastNewLine() Section {
	newLineRegex := MustCompile(`\n`)

	fileFromSection := Section{ s.end, len(s.source), s.source }

	var lastNewLine  Section
	var err error

	lastNewLine, err = fileFromSection.find(newLineRegex)

	//no more new lines in the string. so set to end of string.
	if err != nil {
		lastNewLine = Section{ len(s.source), len(s.source), s.source }
	}

	return lastNewLine
}

func (s *Section) getLines() []Section {

	everythingButNewLineRegex := MustCompile(`[^\n]+`)

	firstNewLine := s.getFirstNewLine()
	lastNewLine := s.getLastNewLine()

	sectionAsLines := Section{firstNewLine.start, lastNewLine.end, s.source}

	return sectionAsLines.findAll(everythingButNewLineRegex)
}


func (s *Section) findFirstCodeBlock(blockPattern *CodeBlockPattern) (Section, error) {
	matches := s.findAll(blockPattern.whole)
	blockStack := stack.New()

	for _, match := range matches {

		symbol := match.toString()

		if blockPattern.isOpen(symbol) {

			blockStack.Push(match)

		} else if blockPattern.isClosed(symbol) {

			stackLength := blockStack.Len()

			if stackLength > 1 {
				blockStack.Pop()
			} else if stackLength == 1 {
				pop := blockStack.Pop()

				startMatch, ok := pop.(Section)
				if ok {
					return Section{startMatch.start, match.end, startMatch.source}, nil
				}

			}

		}

	}

	return Section{}, errors.New("no matches found in this section")
}

func main() {
	re := MustCompile(`ab?`)
	fmt.Println(re.FindStringIndex("tablett"))
	fmt.Println(re.FindStringIndex("foo") == nil)
}
