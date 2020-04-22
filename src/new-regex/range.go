package new_regex

import (
	"math"
	"strings"
)

func Range(exp Expression2, min int, max int) Expression2 {
	return func(iter *Iterator) MatchTree {
		mt := consecutiveCount(iter, exp)
		if isCountRightNumberOfCharacters(len(mt.Children), min, max) {
			mt.isValid = true
			createValue(&mt)
		}else{
			mt.isValid = false
			cullChildren(&mt, max)
			createValue(&mt)
		}
		return mt
	}
}

func cullChildren(mt *MatchTree, max int){
	newChildren := []MatchTree{}
	loopLength := int(math.Min(float64(len(mt.Children)),float64(max)))
	for i := 0; i < loopLength; i++ {
		newChildren = append(newChildren, mt.Children[i])
	}
	mt.Children = newChildren
}

func createValue(mt *MatchTree){
	sb := strings.Builder{}
	for _, child := range mt.Children {
		sb.WriteString(child.Value)
	}
	mt.Value = sb.String()
}

func consecutiveCount(iter *Iterator, exp Expression2) MatchTree {
	mt := MatchTree{}
	for iter.HasNext() {
		startingIndex := iter.GetIndex()
		match := exp(iter)
		if match.isValid{
			mt.Children = append(mt.Children, match) //I think children need to be pointers
		}else{
			iter.Reset(startingIndex) //I think this needs to be tested
			break
		}
	}
	return mt
}

func isCountRightNumberOfCharacters(count, min, max int) bool {
	if max < 0 { //if unlimited
		return count >= min
	}
	return count >= min && count <= max
}
