package main

import (
	"fmt"
	"testing"
)

func Test_createSectionFromString(t *testing.T) {

	original := "Gerald"

	section := createSectionFromString(original)

	fmt.Println("createSectionFromString:")

	if section.source != original {
		t.Errorf("\tSource does not match the original string. got: %s, want: %s", section.source, original)
	}else{
		fmt.Printf("\tsource matches original string: %s vs %s\n", section.source, original)
	}

	if section.start != 0 {
		t.Errorf("\tStart is not correct. got: %d, want: %d", section.start, 0)
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
		fmt.Printf("\ttoString() matches original string: %s vs %s\n", section.source, original)
	}

}