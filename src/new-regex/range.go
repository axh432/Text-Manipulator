package new_regex

import (
	"strings"
)

func Range(exp Expression, min int, max int) Expression {
	return func(iter *Iterator) MatchTree {
		mt := MatchTree{}
		mt.Label = "Range"
		consecutiveCount(iter, exp, &mt)
		createValue(&mt)
		count := len(mt.Children)

		if countAboveMax(count, max){
			mt.isValid = false
			mt.DebugLine = mt.Value + " <- there are more than the max value of " + string(max)
			return mt
		}

		if countBelowMin(count, min){
			mt.isValid = false
			mt.DebugLine = mt.Value + " <- there are less than the min value of " + string(min)
			return mt
		}

		mt.isValid = true
		return mt
	}
}

func createValue(mt *MatchTree){
	sb := strings.Builder{}
	for _, child := range mt.Children {
		sb.WriteString(child.Value)
	}
	mt.Value = sb.String()
}

func consecutiveCount(iter *Iterator, exp Expression, mt *MatchTree) {
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
}

func countAboveMax(count, max int) bool {
	if max < 0 {
		return false //negative numbers classed as infinity
	}else{
		return count > max
	}
}

func countBelowMin(count, min int) bool {
	return count < min
}
