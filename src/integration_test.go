package main

import (
	"fmt"
	. "regexp"
	"sort"
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

	exampleBuild := fileAsSection.find(MustCompile(`Example Build`))

	if exampleBuild.isEmpty() {
		t.Errorf("Example build could not be found. Did you forget to QuoteMeta?")
		return
	}
   
	fmt.Println(exampleBuild.toString())

	restOfFile := Section{ exampleBuild.start, fileAsSection.end, fileAsSection.source }
 
	fmt.Println(restOfFile.toString())
	
	codeBlockPattern := newOpenClosedPatternRegex(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock := restOfFile.findFirstCodeBlock(codeBlockPattern)

	if codeBlock.isEmpty() {
		t.Errorf("Code block could not be found.")
		return
	}

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

	exampleBuild := fileAsSection.find(MustCompile(QuoteMeta(`stage('Example Build') {`)))

	if exampleBuild.isEmpty() {
		t.Errorf("Example build could not be found. Did you forget to QuoteMeta?")
		return
	}

	restOfFile := Section{ exampleBuild.start, fileAsSection.end, fileAsSection.source }

	codeBlockPattern := newOpenClosedPatternRegex(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock := restOfFile.findFirstCodeBlock(codeBlockPattern)

	if codeBlock.isEmpty() {
		t.Errorf("Code block could not be found.")
		return
	}

	lineMatches := codeBlock.findLinesContaining(MustCompile(`sh 'mvn`))

	edit := Edit{ lineMatches[0], "\t\t\t\tWell you can just go and get the hell on" }

	editQueue := EditQueue{}

	editQueue.Add(edit)

	fmt.Printf("%s\n", editQueue.ApplyEdits())

}

//this edit is both the last and the start
func Test_findAndModifySection_replace_start(t *testing.T){

	edit := Edit{ section: Section{ 0, 0, codeSection }, replace:"This should not remove" }

	editQueue := EditQueue{}

	editQueue.Add(edit)

	fmt.Printf("\n%s\n", editQueue.ApplyEdits())

}

func Test_findAndModifySection_replace_end(t *testing.T){

	edit := Edit{ section: Section{ len(codeSection), len(codeSection), codeSection }, replace:"This should not remove" }

	editQueue := EditQueue{}

	editQueue.Add(edit)

	fmt.Printf("\n%s\n", editQueue.ApplyEdits())

}

func Test_findAndModifySection_multiple_replace(t *testing.T){
	edit := Edit{ section: Section{ 0, 0, codeSection }, replace:"This should not remove" }

	edit2 := Edit{ section: Section{ len(codeSection), len(codeSection), codeSection }, replace:"This should not remove" }

	editQueue := EditQueue{}

	editQueue.Add(edit)
	editQueue.Add(edit2)

	fmt.Printf("\n%s\n", editQueue.ApplyEdits())
}

func Test_addADeployStageToJenkinsFile(t *testing.T){

	editQueue := EditQueue{}

	fileAsSection := createSectionFromString(codeSection)

	exampleBuild := fileAsSection.find(MustCompile(QuoteMeta(`stage('Example Build') {`)))

	if exampleBuild.isEmpty() {
		t.Errorf("Example build could not be found. Did you forget to QuoteMeta?")
		return
	}

	codeBlockAndHeader := fitToLines(getHeaderAndCodeBlock(exampleBuild))

	if codeBlockAndHeader.isEmpty() {
		t.Errorf("codeBlockAndHeader could not be found.")
		return
	}

	replace := createCodeInsert("        ")

	editQueue.Add(insertBeforeEdit(codeBlockAndHeader, replace + "\n"))

	fmt.Printf("%s\n", editQueue.ApplyEdits())

}

func Test_addAnExtraStepToExampleBuild(t *testing.T){

	editQueue := EditQueue{}

	fileAsSection := createSectionFromString(codeSection)

	codeBlockAndHeader := findHeaderAndCodeBlock(`stage('Example Build') {`, fileAsSection)

	if codeBlockAndHeader.isEmpty() {
		t.Errorf("codeBlockAndHeader could not be found.")
		return
	}

	agent := findHeaderAndCodeBlock(`agent`, codeBlockAndHeader)

	body := getBodyOfCodeBlock(agent)

	editQueue.Add(replaceEdit(body, "\n" + body.toString() + "\n" ))

	fmt.Printf("%s\n", editQueue.ApplyEdits())

}

func replaceEdit(section Section, replacement string) Edit {
	return Edit{ section, replacement }
}

func deleteEdit(section Section) Edit {
	return Edit{ section, "" }
}

func insertBeforeEdit(section Section, toBeInserted string) Edit {
	return Edit{ section, toBeInserted + section.toString() }
}

func insertAfterEdit(section Section, toBeInserted string) Edit {
	return Edit{ section, section.toString() + toBeInserted }
}

func createCodeInsert(firstIndent string) string {
	var sb Builder

	sb.WriteString("\n" + firstIndent + "stage('Example Deploy') {")
	sb.WriteString("\n" + firstIndent + "    echo 'Hello, Deploy'")
	sb.WriteString("\n" + firstIndent + "}")

	return sb.String()
}


func fitToLines(section Section) Section {
	lines := section.getLines()

	if len(lines) == 0 {
		return Section{}
	}

	return Section { lines[0].start, lines[len(lines)-1].end, section.source }
}

func getBodyOfCodeBlock(section Section) Section {
	codeBlockPattern := newOpenClosedPatternRegex(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock := section.findFirstCodeBlock(codeBlockPattern)

	if codeBlock.isEmpty() {
		return Section{}
	}

	return Section{ codeBlock.start + 1, codeBlock.end -1, codeBlock.source }
}

func findHeaderAndCodeBlock(headerToFind string, section Section) Section {

	header := section.find(MustCompile(QuoteMeta(headerToFind)))

	if header.isEmpty() {
		return Section{}
	}

	return getHeaderAndCodeBlock(header)
}

func getHeaderAndCodeBlock(header Section) Section {
	restOfFile := Section{ header.start, len(header.source), header.source }

	codeBlockPattern := newOpenClosedPatternRegex(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock := restOfFile.findFirstCodeBlock(codeBlockPattern)

	if codeBlock.isEmpty() {
		return Section{}
	}

	return Section{ header.start, codeBlock.end, codeBlock.source }
}

func applyFormatting(line string, formatting string) string {
	return formatting + line
}

func getIndentationFromLine(line Section) string {
	spaceRegex := MustCompile(`\A\s+`)
	lineAsString := line.toString()
	return spaceRegex.FindString(lineAsString)
}

//this doesnt handle tabs or other types of whitespace at the moment
func determineFormatting(input *Section) []string {

	lines := input.getLines()
	spaceRegex := MustCompile(`\A +`)
	indentSlice := []string{}
	indentMap := map[int]string{}

	for _, line := range lines {

		lineAsString := line.toString()
		leadingWhitespace := spaceRegex.FindString(lineAsString)

		lengthOfWhitespace := len(leadingWhitespace)

		if lengthOfWhitespace > 0 {
			indentMap[lengthOfWhitespace] = leadingWhitespace
		}

	}

	for _, value := range indentMap {
		indentSlice = append(indentSlice, value)
	}

	sort.Slice(indentSlice, func(i, j int) bool {
		return len(indentSlice[i]) < len(indentSlice[j])
	})

	return indentSlice
}

func printOutStringInfo(input string) {
	fmt.Printf("length of string: %d\n", len(input))
	fmt.Println(getStringAsMetaCharacters(input))
}

func getStringAsMetaCharacters(input string) string {
	return fmt.Sprintf("%q", ReplaceAll(input, " ", "█"))
}