package new_regex

func SequenceOfCharacters(sequence string) Expression {
	return func(iter *Iterator) bool {

		if sequence == "" {
			return false
		}

		for _, char := range sequence {
			if !iter.HasNext() || char != iter.Next() {
				return false
			}
		}
		return true
	}
}
