package new_regex

import "strings"

func SequenceOfCharacters(sequence string) Expression {
	return func(iter *Iterator) MatchTree {

		if sequence == "" {
			return MatchTree{}
		}

		sb := strings.Builder{}
		for _, char := range sequence {
			if !iter.HasNext() || char != iter.Next() {
				return MatchTree{
					isValid:  false,
					Value:    sb.String(),
					Label:    "",
					Children: nil,
				}
			}else{
				sb.WriteRune(char)
			}
		}

		return MatchTree{
			isValid:  true,
			Value:    sb.String(),
			Label:    "",
			Children: nil,
		}
	}
}
