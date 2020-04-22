package new_regex

import "strings"

type MatchTree struct {
	isValid  bool
	Value    string
	Name     string
	Children []MatchTree
}

func (mt *MatchTree) toString() string {
	sb := strings.Builder{}
	toStringRecursive(mt, &sb, 0)
	return sb.String()
}

func toStringRecursive(mt *MatchTree, sb *strings.Builder, level int){
	for i := 0; i < level; i++ {
		sb.WriteString("\t")
	}
	sb.WriteString("<-[")
	sb.WriteString(mt.Value)
	sb.WriteString("]->")
	sb.WriteString("\n")
	for _, child := range mt.Children {
		toStringRecursive(&child, sb, level + 1)
	}
}

type Expression func(iter *Iterator) bool

type Expression2 func(iter *Iterator) MatchTree
