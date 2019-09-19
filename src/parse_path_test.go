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

	query := "/CurlyCodeBlockWithHeader[stage('Example Build')]/CurlyCodeBlockWithHeader[steps]/Line[1]"


}

func runQuery(query string, section Section) ([]Section, error) {

	commands := strings.Split(query, "/")

	return executeCommands(commands, 0, section)

}

func executeCommands(commands []string, index int, section Section) ([]Section, error) {

	var grammaticalExpName string
	var searchString string
	var sectionsToReturn []Section
	var sectionsFromGrammaticalExp []Section


	command := commands[index]

	if !strings.Contains(command, "[") {
		return []Section{}, errors.New("given command does not contain [] with a search query, " + command)
	}

	grammaticalExpAndSearchStr := strings.Split(command, "[")

	grammaticalExpName = grammaticalExpAndSearchStr[0]
	searchString = strings.Replace(grammaticalExpAndSearchStr[1], "]", "", 1)

	grammaticalExp := grammar[grammaticalExpName]

	switch grammaticalExp.GetType() {

		case "Composition":
			composition, ok := grammaticalExp.(Composition)

			if !ok {
				return []Section{}, errors.New("Composition GrammaticalExpression could not be casted to Composition")
			}
			sectionsFromGrammaticalExp = executeCompositionExpression(composition, searchString)

		case "Or":
			fmt.Println("Linux.")
		case "SinglePattern":
			fmt.Println("Linux.")
		case "MatchedSinglePattern":
			fmt.Println("OS X.")
		case "OpenClosedPattern":
			fmt.Println("Linux.")
		default:

	}

	if index == len(commands)-1 {
		return sectionsFromGrammaticalExp, nil
	}

	for _, sectionFromGrammaticalExp := range sectionsFromGrammaticalExp {

		sectionsFromNextCommand, err := executeCommands(commands, index + 1, sectionFromGrammaticalExp)

		if err != nil {
			return []Section{}, err
		}

		sectionsToReturn = append(sectionsToReturn, sectionsFromNextCommand...)
	}

	return sectionsToReturn, nil
}

func executeCompositionExpression(comp Composition, query string) []Section {
	return []Section{}
}

