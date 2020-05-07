package new_regex

import (
	"fmt"
	"strings"
)

type MatchTree struct {
	IsValid   bool
	Value     string
	Type      string
	Label     string
	Children  []MatchTree
	DebugLine string
}

type MatchTreeVisitor func(mt *MatchTree)

type TypeCounter struct {
	setOfCharsCount int
	seqOfCharsCount int
	sequenceCount   int
	setCount        int
	rangeCount      int
	labelCount      int
}

func (mt *MatchTree) acceptVisitor(visit MatchTreeVisitor) {
	visit(mt)
	for _, child := range mt.Children {
		child.acceptVisitor(visit)
	}
}

func (mt MatchTree) toMermaidDiagram() string {
	counter := TypeCounter{}
	definitions := strings.Builder{}
	links := strings.Builder{}

	links.WriteString("classDiagram")

	toMermaidDiagramRecursive(&mt, "", &counter, &links, &definitions)

	return fmt.Sprintf("%s\n%s", links.String(), definitions.String())
}

func toMermaidDiagramRecursive(mt *MatchTree, parentName string, counter *TypeCounter, links *strings.Builder, definitions *strings.Builder) {

	var name string

	switch mt.Type {
	case "SequenceOfCharacters":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.seqOfCharsCount)
		counter.seqOfCharsCount++
		break
	case "SetOfCharacters":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.setOfCharsCount)
		counter.setOfCharsCount++
		break
	case "Set":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.setCount)
		counter.setCount++
		break
	case "Range":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.rangeCount)
		counter.rangeCount++
		break
	case "Sequence":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.sequenceCount)
		counter.sequenceCount++
		break
	case "Label":
		name = fmt.Sprintf("%s_%d", mt.Type, counter.labelCount)
		counter.labelCount++
		break
	}

	if parentName != "" {
		links.WriteString(fmt.Sprintf("\n\t%s-->%s", parentName, name))
	}

	classDef := `
class %s {
	IsValid: %t
	Value: "%s"
	Type: %s
	Label: %s
	DebugLine: %s
}`

	definitions.WriteString(fmt.Sprintf(classDef, name, mt.IsValid, mt.Value, mt.Type, mt.Label, mt.DebugLine))

	for _, child := range mt.Children {
		toMermaidDiagramRecursive(&child, name, counter, links, definitions)
	}
}

func (mt *MatchTree) toString() string {
	sb := strings.Builder{}
	toStringRecursive(mt, &sb, "")
	return sb.String()
}

func toStringRecursive(mt *MatchTree, sb *strings.Builder, levelPadding string) {
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

func validMatchTree(value string, Type string, children []MatchTree) MatchTree {
	return MatchTree{
		IsValid:   true,
		Value:     value,
		Type:      Type,
		Label:     "",
		Children:  children,
		DebugLine: "",
	}
}

func invalidMatchTree(value string, Type string, debugLine string) MatchTree {
	return MatchTree{
		IsValid:   false,
		Value:     value,
		Type:      Type,
		Label:     "",
		Children:  nil,
		DebugLine: debugLine,
	}
}
