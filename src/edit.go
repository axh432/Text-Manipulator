package main

import (
	"sort"
	"strings"
)

type Edit struct {
	section Section
	replace string
}

type EditQueue struct {
	Edits []Edit
}

func (eq *EditQueue) Add(edit Edit) {

	eq.Edits = append(eq.Edits, edit)

	sort.Slice(eq.Edits, func(i, j int) bool {
		return eq.Edits[i].section.start < eq.Edits[j].section.start
	})
}

func (eq *EditQueue) ApplyEdits() string {

	var sb strings.Builder

	for index, currentEdit := range eq.Edits {

		//fill in the gap between previous and current segment
		var gapStart int

		if index == 0 {
			gapStart = 0
		}else{
			previousEdit := eq.Edits[index -1]
			gapStart = previousEdit.section.end //you can take the previous section end as a start and it wont be repeated.
		}

		if gapStart != currentEdit.section.start {
			sb.WriteString(currentEdit.section.source[gapStart:currentEdit.section.start])
		}

		//write the current segment
		sb.WriteString(currentEdit.replace)

	}

	lastEdit := eq.Edits[len(eq.Edits)-1]
	sourceEnd := len(lastEdit.section.source)

	if lastEdit.section.end < sourceEnd {
		sb.WriteString(lastEdit.section.source[lastEdit.section.end:sourceEnd])
	}

	return sb.String()
}

