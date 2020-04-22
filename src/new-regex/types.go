package new_regex

import "strings"

type MatchTree struct {
	isValid  bool
	Value    string
	Label    string
	Children []MatchTree
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

type Expression func(iter *Iterator) MatchTree
