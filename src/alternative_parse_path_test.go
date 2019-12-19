package main

import (
    "errors"
    "fmt"
    "regexp"
    "testing"
)

func Test_parsePath(t *testing.T) {

    altQuery := `/"Example Build')"|CurlyCodeBlock/agent|CurlyCodeBlock/maven`

    re := regexp.MustCompile(`[|\/]`)

    splitStrings := splitKeepDelimiter(re, altQuery)

    section, err := executeCommands(splitStrings, createSectionFromString(codeSection))

    if err != nil {
        t.Error(err)
    }else{
        fmt.Printf(section.toString())
    }
}

func executeCommands(splitStrings []string, section Section) (Section, error) {

    if len(splitStrings) == 0 {
        return section, nil
    }

    command := splitStrings[0]
    search := splitStrings[1]

    newSection, err := executeCommand(command, search, section)

    if err != nil {
        return newSection, err
    }

    return executeCommands(splitStrings[2:], newSection)
}

func executeCommand(command string, search string, section Section) (Section, error) {

    if command == "/" {
        return searchWithin(search, section)
    } else if command == "|" {
        return searchAdjacent(search, section)
    }

    return Section{}, errors.New("Unknown command type: " + command)
}

func searchWithin(search string, section Section) (Section, error) {

    if search == "CurlyCodeBlock" {
        codeBlockPattern := newOpenClosedPatternRegex(regexp.MustCompile("[{}]"), regexp.MustCompile("{"), regexp.MustCompile("}"))
        codeBlock := section.findFirstCodeBlock(codeBlockPattern)
        return checkForError(codeBlock, errors.New("Couldn't find a CurlyCodeBlock inside: " + section.toString()))
    } else if search == "Line" {
        lines := section.getLines()
        if len(lines) == 0 {
            return Section{}, errors.New("Couldn't retrieve the line for: " + section.toString())
        }
        return lines[0], nil
    }

    search = removeOuterQuotesFromSearch(search)
    newSection := section.find(regexp.MustCompile(regexp.QuoteMeta(search)))

    return checkForError(newSection, errors.New("Couldn't find: " + search + ", within: " + section.toString()))
}

func searchAdjacent(search string, section Section) (Section, error) {

    restOfFile := Section{ section.end, len(section.source), section.source }

    if search == "CurlyCodeBlock" {
        codeBlockPattern := newOpenClosedPatternRegex(regexp.MustCompile("[{}]"), regexp.MustCompile("{"), regexp.MustCompile("}"))
        codeBlock := restOfFile.findFirstCodeBlock(codeBlockPattern)

        err := errors.New("Couldn't find a CurlyCodeBlock adjacent to: " + section.toString())

        if codeBlock.isEmpty() {
            return codeBlock, err
        }

        if !sectionsAreAdjacent(section, codeBlock) {
            return codeBlock, err
        }

        return codeBlock, nil
    }

    if search == "Line" {
        lines := section.getLines()
        if len(lines) == 0 {
            return Section{}, errors.New("Couldn't retrieve the line for: " + section.toString())
        }
        return lines[0], nil
    }

    search = removeOuterQuotesFromSearch(search)
    newSection := restOfFile.find(regexp.MustCompile(regexp.QuoteMeta(search)))
    err :=  errors.New("Couldn't find: " + search + ", within: " + section.toString())

    if newSection.isEmpty() {
        return newSection, err
    }

    if !sectionsAreAdjacent(section, newSection) {
        return newSection, err
    }

    return newSection, nil
}

func sectionsAreAdjacent(startSection Section, endSection Section) bool {

    gapBetween := Section{ startSection.end, endSection.start, startSection.source }

    anyNonWhitespace := regexp.MustCompile(`\S`)

    if(anyNonWhitespace.MatchString(gapBetween.toString())){
        return false
    }

    return true
}

func removeOuterQuotesFromSearch(search string) string {

    searchEndIndex := len(search)

    if search[0:1] == `"` && search[searchEndIndex-1:searchEndIndex] == `"` {
        return search[1:searchEndIndex-1]
    }

    return search
}

func checkForError(section Section, err error) (Section, error){

    if section.isEmpty() {
        return section, err
    }

    return section, nil
}
