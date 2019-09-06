package main

import (
	"fmt"
	. "regexp"
	"strings"
	"testing"
)

const blockOfText = `
				//The following is a block of code
				public static void printNerp(String toBeNerped) {

					if(toBeNerped.contains("nerp")){
						System.out.println("Nerp!")
					}

					System.out.println("Herp!")

				}
				`

func Test_createSectionFromString(t *testing.T) {

	original := "Gerald"

	section := createSectionFromString(original)

	fmt.Println("createSectionFromString:")

	if section.source != original {
		t.Errorf("\tsource does not match the original string. got: %v, want: %v", section.source, original)
	}else{
		fmt.Printf("\tsource matches original string: %s vs %s\n", section.source, original)
	}

	if section.start != 0 {
		t.Errorf("\tStart is not correct. got: %v, want: %v", section.start, 0)
	}else{
		fmt.Printf("\tstart is correct: %d vs %d\n", section.start, 0)
	}

	if section.end != len(original) {
		t.Errorf("\tEnd is not correct. got: %d, want: %d", section.end, len(original))
	}else{
		fmt.Printf("\tend is correct %d vs %d\n", section.end, len(original))
	}

	if section.toString() != original {
		t.Errorf("\ttoString() is not correct. got: %s, want: %s", section.toString(), original)
	}else{
		fmt.Printf("\ttoString() matches original string: %s vs %s\n", section.toString(), original)
	}

}

func Test_createSectionFromMatch(t *testing.T) {

	original := "Gerald"

	pattern := MustCompile(`rald`)

	match := pattern.FindStringIndex(original)

	section := createSectionFromMatch(match, 0, original)

	fmt.Println("createSectionFromMatch:")

	if section.source != original {
		t.Errorf("\tsource does not match the original string. got: %v, want: %v", section.source, original)
	}else{
		fmt.Printf("\tsource matches original string: %s vs %s\n", section.source, original)
	}

	if section.start != 2 {
		t.Errorf("\tStart is not correct. got: %v, want: %v", section.start, 2)
	}else{
		fmt.Printf("\tstart is correct: %d vs %d\n", section.start, 2)
	}

	if section.end != 6 {
		t.Errorf("\tEnd is not correct. got: %d, want: %d", section.end, 6)
	}else{
		fmt.Printf("\tend is correct %d vs %d\n", section.end, 6)
	}

	if section.toString() != "rald" {
		t.Errorf("\ttoString() is not correct. got: %s, want: %s", section.toString(), "rald")
	}else{
		fmt.Printf("\ttoString() matches original string: %s vs %s\n", section.toString(), "rald")
	}

}

func Test_createSubSectionFromMatch(t *testing.T) {

	original := "moreGerald"

	subOriginal := original[4:10]

	pattern := MustCompile(`rald`)

	match := pattern.FindStringIndex(subOriginal)

	section := createSectionFromMatch(match, 4, original)

	fmt.Println("createSubSectionFromMatch:")

	if section.source != original {
		t.Errorf("\tsource does not match the original string. got: %v, want: %v", section.source, original)
	}else{
		fmt.Printf("\tsource matches original string: %s vs %s\n", section.source, original)
	}

	if section.start != 6 {
		t.Errorf("\tStart is not correct. got: %v, want: %v", section.start, 6)
	}else{
		fmt.Printf("\tstart is correct: %d vs %d\n", section.start, 6)
	}

	if section.end != 10 {
		t.Errorf("\tEnd is not correct. got: %d, want: %d", section.end, 10)
	}else{
		fmt.Printf("\tend is correct %d vs %d\n", section.end, 10)
	}

	if section.toString() != "rald" {
		t.Errorf("\ttoString() is not correct. got: %s, want: %s", section.toString(), "rald")
	}else{
		fmt.Printf("\ttoString() matches original string: %s vs %s\n", section.toString(), "rald")
	}

}

func Test_find(t *testing.T) {

	fmt.Println("find:")

	original := "Gerald"

	section := createSectionFromString(original)

	foundSection, err := section.find(MustCompile(`rald`))

	if err != nil {
		t.Errorf(err.Error())
	}

	if foundSection.source != original {
		t.Errorf("\tsource does not match the original string. got: %v, want: %v", foundSection.source, original)
	}else{
		fmt.Printf("\tsource matches original string: %s vs %s\n", foundSection.source, original)
	}

	if foundSection.start != 2 {
		t.Errorf("\tStart is not correct. got: %v, want: %v", foundSection.start, 2)
	}else{
		fmt.Printf("\tstart is correct: %d vs %d\n", foundSection.start, 2)
	}

	if foundSection.end != 6 {
		t.Errorf("\tEnd is not correct. got: %d, want: %d", foundSection.end, 6)
	}else{
		fmt.Printf("\tend is correct %d vs %d\n", foundSection.end, 6)
	}

	if foundSection.toString() != "rald" {
		t.Errorf("\ttoString() is not correct. got: %s, want: %s", foundSection.toString(), "rald")
	}else{
		fmt.Printf("\ttoString() matches substring string: %s vs %s\n", foundSection.toString(), "rald")
	}

}

func Test_findMultiLine(t *testing.T) {

	fmt.Println("findMultiline:")

	original := `something
					Gerald
				this way comes`

	section := createSectionFromString(original)

	foundSection, err := section.find(MustCompile(`rald`))

	if err != nil {
		t.Errorf(err.Error())
	}

	if foundSection.source != original {
		t.Errorf("\tsource does not match the original string. got: %v, want: %v", foundSection.source, original)
	}else{
		fmt.Printf("\tsource matches original string: %s vs %s\n", foundSection.source, original)
	}

	if foundSection.start != 17 {
		t.Errorf("\tStart is not correct. got: %v, want: %v", foundSection.start, 17)
	}else{
		fmt.Printf("\tstart is correct: %d vs %d\n", foundSection.start, 17)
	}

	if foundSection.end != 21 {
		t.Errorf("\tEnd is not correct. got: %d, want: %d", foundSection.end, 21)
	}else{
		fmt.Printf("\tend is correct %d vs %d\n", foundSection.end, 21)
	}

	if foundSection.toString() != "rald" {
		t.Errorf("\ttoString() is not correct. got: %s, want: %s", foundSection.toString(), "rald")
	}else{
		fmt.Printf("\ttoString() matches substring string: %s vs %s\n", foundSection.toString(), "rald")
	}

}

func Test_findWithinSubSection(t *testing.T) {

	fmt.Println("findWithinSubSection:")

	original := "moreGerald"

	section := createSectionFromString(original)

	foundSection, err := section.find(MustCompile(`Gerald`))

	if err != nil {
		t.Errorf(err.Error())
	}

	foundSubSection, subErr := foundSection.find(MustCompile(`rald`))

	if subErr != nil {
		t.Errorf(subErr.Error())
	}

	if foundSubSection.source != original {
		t.Errorf("\tsource does not match the original string. got: %v, want: %v", foundSubSection.source, original)
	}else{
		fmt.Printf("\tsource matches original string: %s vs %s\n", foundSubSection.source, original)
	}

	if foundSubSection.start != 6 {
		t.Errorf("\tStart is not correct. got: %v, want: %v", foundSubSection.start, 6)
	}else{
		fmt.Printf("\tstart is correct: %d vs %d\n", foundSubSection.start, 6)
	}

	if foundSubSection.end != 10 {
		t.Errorf("\tEnd is not correct. got: %d, want: %d", foundSubSection.end, 10)
	}else{
		fmt.Printf("\tend is correct %d vs %d\n", foundSubSection.end, 10)
	}

	if foundSubSection.toString() != "rald" {
		t.Errorf("\ttoString() is not correct. got: %s, want: %s", foundSubSection.toString(), "rald")
	}else{
		fmt.Printf("\ttoString() matches substring string: %s vs %s\n", foundSubSection.toString(), "rald")
	}

}

func Test_findAll(t *testing.T) {

	fmt.Println("findAll:")

	original := "Geraldralb"

	section := createSectionFromString(original)

	foundSections := section.findAll(MustCompile(`ral[d|b]`))

	if len(foundSections) != 2 {
		t.Errorf("\tthe number of sections found in the string does not match what was expected. got: %d, want: %d", len(foundSections), 2)
	}else{
		fmt.Printf("\tthe number of sections found is correct: %d vs %d\n", len(foundSections), 2)
	}

	for _, foundSection := range foundSections {

		if foundSection.source != original {
			t.Errorf("\tsource does not match the original string. got: %v, want: %v", foundSection.source, original)
		}else{
			fmt.Printf("\tsource matches original string: %s vs %s\n", foundSection.source, original)
		}

		if foundSection.toString() != "rald" && foundSection.toString() != "ralb" {
			t.Errorf("\ttoString() is not correct. got: %s, want: %s", foundSection.toString(), "rald or ralb")
		}else{
			fmt.Printf("\ttoString() matches substring string: %s vs %s\n", foundSection.toString(), "rald or ralb")
		}

	}

}

func Test_findWithMultiple(t *testing.T) {

	fmt.Println("findWithMultiple:")

	original := "Geraldrald"

	section := createSectionFromString(original)

	foundSection, err := section.find(MustCompile(`rald`))

	if err != nil {
		t.Errorf(err.Error())
	}

	if foundSection.source != original {
		t.Errorf("\tsource does not match the original string. got: %v, want: %v", foundSection.source, original)
	}else{
		fmt.Printf("\tsource matches original string: %s vs %s\n", foundSection.source, original)
	}

	if foundSection.start != 2 {
		t.Errorf("\tStart is not correct. got: %v, want: %v", foundSection.start, 2)
	}else{
		fmt.Printf("\tstart is correct: %d vs %d\n", foundSection.start, 2)
	}

	if foundSection.end != 6 {
		t.Errorf("\tEnd is not correct. got: %d, want: %d", foundSection.end, 6)
	}else{
		fmt.Printf("\tend is correct %d vs %d\n", foundSection.end, 6)
	}

	if foundSection.toString() != "rald" {
		t.Errorf("\ttoString() is not correct. got: %s, want: %s", foundSection.toString(), "rald")
	}else{
		fmt.Printf("\ttoString() matches substring string: %s vs %s\n", foundSection.toString(), "rald")
	}

}

func Test_findAllStartEndPattern(t *testing.T) {

	fmt.Println("findAllStartEndPattern:")

	original := `

				Gerald is rad
				Gerald is great
				Gerald is life
				`

	section := createSectionFromString(original)

	foundLines := section.findAllStartEndPattern(MustCompile("Gerald"), MustCompile("\n"))

	for _, foundLine := range foundLines {
		if foundLine.source != original {
			t.Errorf("\tsource does not match the original string. got: %v, want: %v", foundLine.source, original)
		}else{
			fmt.Println("\tsource matches original string")
		}

		if strings.Contains("Gerald", foundLine.toString()){
			t.Errorf("\ttoString() is not correct. got: %s", foundLine.toString())
		}else{
			fmt.Printf("\ttoString() has the correct line: %s", foundLine.toString())
		}

	}

}

func Test_findLinesContaining(t *testing.T) {

	fmt.Println("findLinesContaining:")

	original := `

				Gerald is rad
				Gerald is great
				Gerald is life
				`

	section := createSectionFromString(original)

	foundLines := section.findLinesContaining(MustCompile("Gerald"))

	for _, foundLine := range foundLines {
		if foundLine.source != original {
			t.Errorf("\tsource does not match the original string. got: %v, want: %v", foundLine.source, original)
		}else{
			fmt.Println("\tsource matches original string")
		}

		if strings.Contains("Gerald", foundLine.toString()){
			t.Errorf("\ttoString() is not correct. got: %s", foundLine.toString())
		}else{
			fmt.Printf("\ttoString() has the correct line: %s", foundLine.toString())
		}

	}

}

func Test_findFirstCodeBlock(t *testing.T) {

	fmt.Println("findFirstCodeBlock:")

	original := `
				//The following is a block of code
				public static void printNerp(String toBeNerped) {
					
					if(toBeNerped.contains("nerp")){
						System.out.println("Nerp!")
					}

					System.out.println("Herp!")

				}
				`

	expected := `{
					
					if(toBeNerped.contains("nerp")){
						System.out.println("Nerp!")
					}

					System.out.println("Herp!")

				}`

	section := createSectionFromString(original)

	codeBlockPattern := newCodeBlockPattern(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock, err := section.findFirstCodeBlock(codeBlockPattern)

	if err != nil {
		t.Errorf(err.Error())
	}

	if codeBlock.toString() != expected {
		t.Errorf(`the code block is incorrect. 
							Expected: %s
							Got: %s `, expected, codeBlock.toString())
	}

	fmt.Printf(codeBlock.toString())
}

func Test_getLines(t *testing.T) {
	fmt.Println("getLines():")

	section := createSectionFromString(blockOfText)

	functionHeader, err := section.find(MustCompile(QuoteMeta(`public static void printNerp(String toBeNerped) {`)))

	if err != nil {
		t.Errorf(err.Error())
	}

	restOfFile := Section{ functionHeader.start, section.end, section.source }

	codeBlockPattern := newCodeBlockPattern(MustCompile("[{}]"), MustCompile("{"), MustCompile("}"))

	codeBlock, err := restOfFile.findFirstCodeBlock(codeBlockPattern)

	if err != nil {
		t.Errorf(err.Error())
	}

	headerAndCodeBlock := Section{ functionHeader.start, codeBlock.end, section.source }

	lines := headerAndCodeBlock.getLines()

	for index, line := range lines {
		fmt.Printf("%d: %s\n", index, line.toString())
	}

}

func Test_getLines_one_word(t *testing.T){
	fmt.Println("getLines():")

	section := createSectionFromString(blockOfText)

	void, err := section.find(MustCompile(QuoteMeta(`void`)))

	if err != nil {
		t.Errorf(err.Error())
	}

	lines := void.getLines()

	for index, line := range lines {
		fmt.Printf("%d: %s\n", index, line.toString())
	}


}

func Test_getLines_no_new_lines_at_end(t *testing.T){
	fmt.Println("getLines():")

	section := createSectionFromString("\n\nThis is dave")

	void, err := section.find(MustCompile(QuoteMeta(`This is dave`)))

	if err != nil {
		t.Errorf(err.Error())
	}

	lines := void.getLines()

	for index, line := range lines {
		fmt.Printf("%d: %s\n", index, line.toString())
	}

}

func Test_getLines_no_new_lines_at_start(t *testing.T){
	fmt.Println("getLines():")

	original := "This is dave\n\n"
	expected := "This is dave"

	section := createSectionFromString(original)

	void, err := section.find(MustCompile(expected))

	if err != nil {
		t.Errorf(err.Error())
	}

	lines := void.getLines()

	for index, line := range lines {
		fmt.Printf("%d: %s\n", index, line.toString())
	}

	if lines[0].toString() != expected {
		t.Errorf("\nExpected: %s\nGot: %s ", expected, lines[0].toString())
	}

}

func Test_getLines_no_new_lines(t *testing.T){
	fmt.Println("getLines():")

	original := "This is dave"
	expected := "This is dave"
	sectionToFind := "dave"

	section := createSectionFromString(original)

	void, err := section.find(MustCompile(QuoteMeta(sectionToFind)))

	if err != nil {
		t.Errorf(err.Error())
	}

	lines := void.getLines()

	for index, line := range lines {
		fmt.Printf("%d: %s\n", index, line.toString())
	}

	if lines[0].toString() != expected {
		t.Errorf("\nExpected: %s\nGot: %s ", expected, lines[0].toString())
	}

}