package main

import (
	"errors"
	"fmt"
	. "regexp"
	"strings"
	"testing"
)

var exampleJenkins string = `pipeline {
    agent none 
    stages {
        stage('Example Build') {
            agent { docker 'maven:3-alpine' } 
            steps {
                echo 'Hello, Maven'
                sh 'mvn --version'
            }
        }
        stage('Example Test') {
            agent { docker 'openjdk:8-jre' } 
            steps {
                echo 'Hello, JDK'
                sh 'java -version'
            }
        }
    }
}`

//Types of CodeBlock:
//nested open closed pattern
//matched single pattern with mitigation


var Line = MatchedSinglePattern{ MustCompile("\n"), []string{} }

var DoubleQuoteString = MatchedSinglePattern{ MustCompile(`"`), []string{ `\"` } }

var SingleQuoteString = MatchedSinglePattern{ MustCompile(`'`), []string{ `\'` } }

var CurlyCodeBlock = newOpenClosedPatternString("{", "}" )

var ParenthesesCodeBlock = newOpenClosedPatternString( "(", ")" )

var Word = SinglePattern { MustCompile(`\w`) }

var ParenthesesCodeBlockWithHeader = Composition{ Word, ParenthesesCodeBlock }

var CurlyHeader = Or { Word, ParenthesesCodeBlockWithHeader }

var CurlyCodeBlockWithHeader = Composition{ CurlyHeader, CurlyCodeBlock }

var grammar = map[string] GrammaticalExpression {
	"Line" : Line,
	"DoubleQuoteString" : DoubleQuoteString,
	"SingleQuoteString" : SingleQuoteString,
	"CurlyCodeBlock" : CurlyCodeBlock,
	"ParenthesesCodeBlock" : ParenthesesCodeBlock,
	"Word" : Word,
	"ParenthesesCodeBlockWithHeader" : ParenthesesCodeBlockWithHeader,
	"CurlyHeader" : CurlyHeader,
	"CurlyCodeBlockWithHeader" : CurlyCodeBlockWithHeader,
}


func Test_parsePath(t *testing.T) {

	//todo: grow out algorithm
	query := "/CurlyCodeBlockWithHeader[stage('Example Build')]/CurlyCodeBlockWithHeader[steps]/Line[1]"

	//todo: find adjacent algorithm
	altQuery := "/stage('Example Build')|CurlyCodeBlock/steps|CurlyCodeBlock/sh 'mvn|Line"


}

func runQuery(query string, section Section) ([]Section, error) {

	commands := strings.Split(query, "/")

	return executeCommands(commands, 0, section)

}

func executeCommands(commands []string, index int, section Section) ([]Section, error) {

	var sectionsFromExpression []Section
	var err error

	command := commands[index]

	grammaticalExp, searchQuery, interpretErr := interpretCommand(command)

	if interpretErr != nil {
		return []Section{}, interpretErr
	}

	sectionsFromExpression, err = executeGrammaticalExpression(grammaticalExp, searchQuery, section)

	if err != nil {
		return []Section{}, err
	}

	if index == len(commands)-1 {
		return sectionsFromExpression, nil
	}

	return executeNextCommand(commands, index + 1, sectionsFromExpression)
}

func executeNextCommand(commands []string, index int, sections []Section) ([]Section, error) {

	var sectionsToReturn []Section

	for _, section := range sections {

		sectionsFromNextCommand, nextCommandErr := executeCommands(commands, index, section)

		sectionsToReturn = append(sectionsToReturn, sectionsFromNextCommand...)

		if nextCommandErr != nil {
			return []Section{}, nextCommandErr
		}
	}

	return sectionsToReturn, nil
}

func interpretCommand(command string) (grammaticalExp GrammaticalExpression, query string, err error){

	grammaticalExpName := getGrammaticalExpFromCommand(command)
	query = getQueryFromCommand(command)

	if len(grammaticalExpName) == 0 || len(query) == 0 {
		return nil, "", errors.New("given command does not contain [] with a search query, " + command)
	}

	return grammar[grammaticalExpName], query, nil
}

func getGrammaticalExpFromCommand(command string) string {

	if !strings.Contains(command, "[") {
		return ""
	}

	grammaticalExpAndSearchStr := strings.Split(command, "[")

	return grammaticalExpAndSearchStr[0]

}

func getQueryFromCommand(command string) string {

	if !strings.Contains(command, "[") {
		return ""
	}

	grammaticalExpAndSearchStr := strings.Split(command, "[")

	return strings.Replace(grammaticalExpAndSearchStr[1], "]", "", 1)

}

func executeGrammaticalExpression(grammaticalExp GrammaticalExpression, searchString string, section Section) ([]Section, error) {

	var sectionsFromGrammaticalExp []Section

	switch grammaticalExp.GetType() {

	case "Composition":
		//return executeCompositionExpression(grammaticalExp, searchString, section)
	case "Or":
		//return executeOrExpression(grammaticalExp, searchString, section)
	case "SinglePattern":
		fmt.Println("SinglePattern")
	case "MatchedSinglePattern":
		fmt.Println("MatchedSinglePattern")
	case "OpenClosedPattern":
		openClosedPattern, err := castToOpenClosedPattern(grammaticalExp)
		if err != nil{
			return []Section{}, err
		}
		return executeOpenClosedPatternExpression(openClosedPattern, section)

	default:
		return []Section{}, errors.New(fmt.Sprintf("Unable to determine the type of grammatical expression: %s", grammaticalExp.GetType()))
	}

	return sectionsFromGrammaticalExp, nil
}

//Todo: figure out what you are doing with error handling!
func castToOpenClosedPattern(grammaticalExp GrammaticalExpression) (OpenClosedPattern, error){
	openClosedPattern, ok := grammaticalExp.(OpenClosedPattern)

	if !ok {
		return OpenClosedPattern{}, errors.New("OpenClosedPattern GrammaticalExpression could not be casted to OpenClosedPattern object. Something is very wrong")
	}

	return openClosedPattern, nil
}

func executeOpenClosedPatternExpression(openClosedPattern OpenClosedPattern, section Section) ([]Section, error) {

	restOfFile := Section{ section.start, len(section.source), section.source }

	codeBlock := restOfFile.findFirstCodeBlock(&openClosedPattern)

	if codeBlock.isEmpty() {
		return []Section{}, errors.New("open closed block could not be found")
	}

	if thereIsSomethingBetweenSections(section, codeBlock) {
		return []Section{}, errors.New("an open closed block could not be found immediately after the given section")
	}

	return []Section{ codeBlock }, nil
}

func thereIsSomethingBetweenSections(startSection Section, endSection Section) bool {

	gapBetween := Section{ startSection.end +1, endSection.start -1, startSection.source }

	anyNonWhitespace := MustCompile(`\S`)

	if(anyNonWhitespace.MatchString(gapBetween.toString())){
		return true
	}

	return false
}





