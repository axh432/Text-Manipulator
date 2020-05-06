package new_regex

import "strings"

type MatchTree struct {
	isValid  bool
	Value    string
	Label    string
	Children []MatchTree
	DebugLine string
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
