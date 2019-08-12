package main

import (
	"fmt"
	"testing"
	. "regexp"
)

func Test_findStringIndex(t *testing.T){

	fileAsSection := createSectionFromString(codeSection)

	fmt.Println(fileAsSection.toString())

	exampleBuild, err := fileAsSection.find(MustCompile("stage\\(\\'Example Build\\'\\)"))

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Println(exampleBuild.toString())

}