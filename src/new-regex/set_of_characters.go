package new_regex

func SetOfCharacters(characters string) Expression2 {
	return func(iter *Iterator) MatchTree {
		if !iter.HasNext() {
			return MatchTree{}
		}
		nextRune := iter.Next()
		for _, char := range characters {
			if char == nextRune {
				return MatchTree {
					isValid:  true,
					Value:    string(nextRune),
					Name:     "",
					Children: nil,
				}
			}
		}
		return MatchTree{}
	}
}
