package main

import (
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


type Composition struct {
	first interface{}
	second interface{}
}

type Or struct {
	first interface{}
	second interface{}
}




var Line = MatchedSinglePattern{ MustCompile("\n"), []string{} }

var DoubleQuoteString = MatchedSinglePattern{ MustCompile(`"`), []string{ `\"` } }

var SingleQuoteString = MatchedSinglePattern{ MustCompile(`'`), []string{ `\'` } }

var CurlyCodeBlock = newOpenClosedPatternString("{", "}" )

var ParenthesesCodeBlock = newOpenClosedPatternString( "(", ")" )

var Word = MustCompile(`\w`)

var ParenthesesCodeBlockWithHeader = Composition{ Word, ParenthesesCodeBlock }

var CurlyHeader = Or { Word, ParenthesesCodeBlockWithHeader }

var CurlyCodeBlockWithHeader = Composition{ CurlyHeader, CurlyCodeBlock }

var grammar = map[string] interface{}{
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

func runQuery(query string, section Section) []Section {

	commands := strings.Split(query, "/")

	return executeCommands(commands, 0, section)

}

func executeCommands(commands []string, index int, section Section) []Section {

	command := commands[index]
	var something interface{}
	var stringToSearch string

	if strings.Contains(command, "[") {

		commandAndSearchStr := strings.Split(command, "[")
		commandAndSearchStr[1] = strings.Replace(commandAndSearchStr[1], "]", "", 1)

		something = grammar[commandAndSearchStr[0]]
		stringToSearch = commandAndSearchStr[1]

	}

	somethingAsSinglePattern 		:= castToMatchedSinglePattern(something)
	somethingAsOpenClosedPattern 	:= castToOpenClosedPattern(something)
	somethingAsOr 					:= castToOr(something)
	somethingAsComposition 			:= castToComposition(something)

	if somethingAsSinglePattern != nil {

	}

}

func castToComposition(object interface{}) Composition {
	var composition Composition

	composition, _ = object.(Composition)

	return composition
}

func castToOr(object interface{}) Or {
	var or Or

	or, _ = object.(Or)

	return or
}

func castToMatchedSinglePattern(object interface{}) MatchedSinglePattern {
	var matchedSinglePattern MatchedSinglePattern

	matchedSinglePattern, _ = object.(MatchedSinglePattern)

	return matchedSinglePattern
}

func castToOpenClosedPattern(object interface{}) OpenClosedPattern {
	var openClosedPattern OpenClosedPattern

	openClosedPattern, _ = object.(OpenClosedPattern)

	return openClosedPattern
}