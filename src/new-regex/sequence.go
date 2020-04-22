package new_regex

import "strings"

func Sequence(expressions ...Expression2) Expression2 {
	return func(iter *Iterator) MatchTree {

		if len(expressions) == 0 {
			return MatchTree{}
		}

		mt := MatchTree{}
		sb := strings.Builder{}

		for _, exp := range expressions {
			match := exp(iter)
			if match.isValid {
				sb.WriteString(match.Value)
				mt.Children = append(mt.Children, match)
				mt.Value = sb.String()
			}else{
				return mt
			}
		}

		mt.isValid = true
		return mt
	}
}
