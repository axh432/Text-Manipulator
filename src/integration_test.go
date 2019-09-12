package main

import (
	"fmt"
	. "regexp"
	"testing"
	."strings"
)

var codeSection string = `pipeline {
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

func Test_findAndModifySection(t *testing.T){

	fileAsSection := createSectionFromString(codeSection)

	exampleBuild, err := fileAsSection.find(MustCompile(QuoteMeta(`stage('Example Build') {`)))

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	restOfFile := Section{ exampleBuild.start, fileAsSection.end, fileAsSection.source }

	codeBlockPattern := newCodeBlockPattern(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock, err := restOfFile.findFirstCodeBlock(codeBlockPattern)

	fmt.Println(codeBlock.toString())

	lineMatches := codeBlock.findLinesContaining(MustCompile(`sh 'mvn`))

	edit := Edit{ lineMatches[0], "Well you can just go and get the hell on\n" }

	editQueue := EditQueue{}

	editQueue.Add(edit)

	fmt.Printf("%s\n", editQueue.ApplyEdits())

}

//this edit is both the last and the start
func Test_findAndModifySection_replace_start(t *testing.T){

	edit := Edit{ section: Section{ 0, 0, codeSection }, replace:"This should remove" }

	editQueue := EditQueue{}

	editQueue.Add(edit)

	fmt.Printf("\n%s\n", editQueue.ApplyEdits())

}

func Test_findAndModifySection_replace_end(t *testing.T){

	edit := Edit{ section: Section{ len(codeSection), len(codeSection), codeSection }, replace:"This should remove" }

	editQueue := EditQueue{}

	editQueue.Add(edit)

	fmt.Printf("\n%s\n", editQueue.ApplyEdits())

}

func Test_findAndModifySection_multiple_replace(t *testing.T){
	edit := Edit{ section: Section{ 0, 0, codeSection }, replace:"This should remove" }

	edit2 := Edit{ section: Section{ len(codeSection), len(codeSection), codeSection }, replace:"This should remove" }

	editQueue := EditQueue{}

	editQueue.Add(edit)
	editQueue.Add(edit2)

	fmt.Printf("\n%s\n", editQueue.ApplyEdits())
}