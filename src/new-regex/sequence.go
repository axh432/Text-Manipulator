package new_regex

import "strings"

func Sequence(expressions ...Expression) Expression {
	return func(iter *Iterator) MatchTree {
		mt := MatchTree{}

		if len(expressions) == 0 {
			mt.DebugLine = "sequence of expressions is empty"
			return mt
		}

		sb := strings.Builder{}

		for _, exp := range expressions {
			match := exp(iter)
			if match.isValid {
				sb.WriteString(match.Value)
				mt.Children = append(mt.Children, match)
				mt.Value = sb.String()
			}else{
				mt.DebugLine = sb.String() + "<- expecting " + match.Label + " here"
				return mt
			}
		}

		mt.isValid = true
		return mt
	}
}
