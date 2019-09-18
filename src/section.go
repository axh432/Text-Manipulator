package main

import (
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

func (s *Section) isEmpty() bool {
	if s.start == 0 && s.end == 0 && len(s.source) == 0 {
		return true
	}

	return false
}

func (s *Section) toString() string {
	return s.source[s.start:s.end]
}

func (s *Section) find(pattern *Regexp) Section {
	start := s.start

	sectionAsString := s.toString()

	match := pattern.FindStringIndex(sectionAsString)

	if match == nil {
		return Section{}
	}

	return createSectionFromMatch(match, start, s.source)
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

		endMatch := restOfSection.find(endPattern)

		if endMatch.isEmpty() {
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

	lastNewLine := fileFromSection.find(newLineRegex)

	//no more new lines in the string. so set to end of string.
	if lastNewLine.isEmpty() {
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


func (s *Section) findFirstCodeBlock(blockPattern *OpenClosedPattern) Section {
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
					return Section{startMatch.start, match.end, startMatch.source}
				}

			}

		}

	}

	return Section{}
}

func main() {
	myString := "ThisWillBeMyFinestWork"

	re := MustCompile(`This`)
	re2 := MustCompile(`Will`)
	re3 := MustCompile(`Work`)

	fmt.Println(re.FindStringIndex(myString))
	fmt.Println(re2.FindStringIndex(myString))
	fmt.Println(re3.FindStringIndex(myString))

	fmt.Println(myString[8:18])
}
