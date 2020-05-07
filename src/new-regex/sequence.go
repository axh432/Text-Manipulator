package new_regex

import (
	"fmt"
	"strings"
)

func Sequence(expressions ...Expression) Expression {
	return func(iter *Iterator) MatchTree {
		if len(expressions) == 0 {
			return invalidMatchTree("", "Sequence", "Sequence:[], NoMatch:number of subexpressions is zero")
		}

		sb := strings.Builder{}
		matches := []MatchTree{}

		startingIndex := iter.index

		for _, exp := range expressions {
			match := exp(iter)
			if match.IsValid {
				sb.WriteString(match.Value)
				matches = append(matches, match)
			}else{
				iter.Reset(startingIndex)
				debugLine := "Sequence:[], NoMatch:string does not match given subexpression"
				if match.Label != "" {
					 debugLine = fmt.Sprintf("Sequence:[], NoMatch:string does not match given subexpression: %s", match.Label)
				}
				return invalidMatchTree("", "Sequence", debugLine)
			}
		}

		return validMatchTree(sb.String(), "Sequence", matches)
	}
}
