package new_regex

import (
	"fmt"
	"strings"
)

type MatchTree struct {
	isValid  bool
	Value    string
	Label    string
	Children []MatchTree
	DebugLine string
}

type TypeCounter struct {
	setOfCharsCount int
	seqOfCharsCount int
	sequenceCount int
	setCount int
	rangeCount int
}

func (mt *MatchTree) toMermaidDiagram() string {
	counter := TypeCounter{}
	definitions := strings.Builder{}
	links := strings.Builder{}

	links.WriteString("classDiagram")

	toMermaidDiagramRecursive(mt, "", &counter, &links, &definitions)

	return fmt.Sprintf("%s\n%s", links.String(), definitions.String())
}

func toMermaidDiagramRecursive(mt *MatchTree, parentName string, counter *TypeCounter, links *strings.Builder, definitions *strings.Builder){
	name := fmt.Sprintf("Node%d", counter.sequenceCount)
	counter.sequenceCount++

	if parentName != "" {
		links.WriteString(fmt.Sprintf("\n\t%s-->%s", parentName, name))
	}

	classDef := `
class %s {
	IsValid: %t
	Value: %s
	Label: %s
	DebugLine: %s
}`

	definitions.WriteString(fmt.Sprintf(classDef, name, mt.isValid, mt.Value, mt.Label, mt.DebugLine))

	for _, child := range mt.Children {
		toMermaidDiagramRecursive(&child, name, counter, links, definitions)
	}
}

func (mt *MatchTree) toString() string {
	sb := strings.Builder{}
	toStringRecursive(mt, &sb, "")
	return sb.String()
}

func toStringRecursive(mt *MatchTree, sb *strings.Builder, levelPadding string){
	levelPadding = levelPadding + "\t\t"
	sb.WriteString(levelPadding)
	sb.WriteString("|")
	sb.WriteString("\n")

	sb.WriteString(levelPadding)
	sb.WriteString("->[")
	if mt.Label != "" {
		sb.WriteString(mt.Label)
		sb.WriteString(":")
	}
	sb.WriteString(mt.Value)
	sb.WriteString("]")
	sb.WriteString("\n")
	for _, child := range mt.Children {
		toStringRecursive(&child, sb, levelPadding)
	}
}

func validMatchTree(value string, children []MatchTree) MatchTree {
	return MatchTree{
		isValid:   true,
		Value:     value,
		Label:     "",
		Children:  children,
		DebugLine: "",
	}
}

func invalidMatchTree(value string, debugline string) MatchTree {
	mt := MatchTree{}
	mt.isValid = false
	mt.Value = value
	mt.DebugLine = debugline
	return mt
}
