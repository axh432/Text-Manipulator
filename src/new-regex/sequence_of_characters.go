package new_regex

import "strings"

func SequenceOfCharacters(sequence string) Expression {
	return func(iter *Iterator) MatchTree {
		mt := MatchTree{}

		if sequence == "" {
			mt.DebugLine = "SequenceOfCharacters:[" + sequence + "], NoMatch:sequence of characters is empty"
			return mt
		}

		sb := strings.Builder{}

		for _, char := range sequence {

			if !iter.HasNext() {
				mt.isValid = false
				mt.Value = sb.String()
				mt.DebugLine = "SequenceOfCharacters:[" + sequence + "], NoMatch:reached end of string before finished"
				return mt
			}

			if char != iter.Next() {
				mt.isValid = false
				mt.Value = sb.String()
				mt.DebugLine = "SequenceOfCharacters:[" + sequence + "], NoMatch: '" + string(char) + "' does not match the sequence"
				return mt
			}

			sb.WriteRune(char)
		}

		mt.isValid = true
		mt.Value = sb.String()
		return mt
	}
}
