package main

import (
	"sort"
	"strings"
)

type Edit struct {
	section *Section
	replace string
}

type EditQueue struct {
	Edits []Edit
}

func (eq *EditQueue) Add(edit Edit) {

	eq.Edits = append(eq.Edits, edit)

	sort.Slice(eq.Edits, func(i, j int) bool {
		return eq.Edits[i].section.start > eq.Edits[j].section.start
	})
}

func (eq *EditQueue) ApplyEdits() string {

	var sb strings.Builder

	for index, edit := range eq.Edits {

		if(index == 0){
			if(edit.section.start > 0){
				sb.WriteString(edit.section.source[0:edit.section.start-1])
				sb.WriteString(edit.replace)
			}
		} else {

			previousEdit := eq.Edits[index - 1]
			gapBetweenSections := edit.section.start - previousEdit.section.end

			if(gapBetweenSections > 0){
				source := edit.section.source
				gapBegin := previousEdit.section.end + 1
				gapEnd := edit.section.start -1

				sb.WriteString(source[gapBegin:gapEnd])
			}

			sb.WriteString(previousEdit.replace)
		}

	}

	return sb.String()
}

