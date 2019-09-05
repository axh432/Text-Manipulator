package main

import (
	"fmt"
	. "regexp"
	"testing"
	."strings"
)

var codeSection string = `
pipeline {
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
}
`

func Test_findAndModifySection(t *testing.T){

	fileAsSection := createSectionFromString(codeSection)

	fmt.Println(fileAsSection.toString())

	exampleBuild, err := fileAsSection.find(MustCompile(`Example Build`))

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Println(exampleBuild.toString())

	restOfFile := Section{ exampleBuild.start, fileAsSection.end, fileAsSection.source }

	fmt.Println(restOfFile.toString())

	codeBlockPattern := newCodeBlockPattern(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock, err := restOfFile.findFirstCodeBlock(codeBlockPattern)

	fmt.Println(codeBlock.toString())

	lineMatches := codeBlock.findLinesContaining(MustCompile(`sh 'mvn`))

	for _, line := range lineMatches {

		if !Contains(line.toString(),`sh 'mvn --version'`) {
			t.Errorf(`expected: sh 'mvn --version', got: %s`, line.toString())
		}else{
			fmt.Println(line.toString())
		}

	}

}

func Test_findSectionWithinGroovyFile(t *testing.T){

	fileAsSection := createSectionFromString(codeSection)

	fmt.Println(fileAsSection.toString())

	exampleBuild, err := fileAsSection.find(MustCompile(`Example Build`))

	if err != nil {
		t.Errorf(err.Error())
		return
	}
   
	fmt.Println(exampleBuild.toString())

	restOfFile := Section{ exampleBuild.start, fileAsSection.end, fileAsSection.source }
 
	fmt.Println(restOfFile.toString())
	
	codeBlockPattern := newCodeBlockPattern(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock, err := restOfFile.findFirstCodeBlock(codeBlockPattern)

	fmt.Println(codeBlock.toString())

	lineMatches := codeBlock.findLinesContaining(MustCompile(`sh 'mvn`))
	
	for _, line := range lineMatches {

		if !Contains(line.toString(),`sh 'mvn --version'`) {
			t.Errorf(`expected: sh 'mvn --version', got: %s`, line.toString())
		}else{
			fmt.Println(line.toString())
		}

	}

}