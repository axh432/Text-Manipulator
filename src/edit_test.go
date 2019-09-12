package main

import (
	"fmt"
	"testing"
)

func Test_editQueueAdd(t *testing.T){

	editQueue := EditQueue{}

	edit1 := Edit{
		Section{ 20, 30, "" },
		"",
	}

	edit2 := Edit{
		Section{ 10, 25, "" },
		"",
	}

	edit3 := Edit{
		Section{ 5, 50, "" },
		"",
	}

	edit4 := Edit{
		Section{ 51, 80, "" },
		"",
	}

	editQueue.Add(edit1)
	editQueue.Add(edit2)
	editQueue.Add(edit3)
	editQueue.Add(edit4)

	for index, edit := range editQueue.Edits {

		if index != 0 {

			previousEdit := editQueue.Edits[index-1]

			if edit.section.start <= previousEdit.section.start {
				t.Errorf(`edits are supposed to be in ascending start index orders`)
				break
			}

		}
		fmt.Printf("%+v\n", edit)

	}

}