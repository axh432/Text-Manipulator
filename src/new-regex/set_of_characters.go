package new_regex

func SetOfCharacters(characters string) Expression {
	return func(iter *Iterator) MatchTree {
		if !iter.HasNext() {
			return invalidMatchTree("", "SetOfCharacters:[" + characters + "], NoMatch:reached end of string before finished")
		}

		startingIndex := iter.index
		nextRune := iter.Next()
		for _, char := range characters {
			if char == nextRune {
				return validMatchTree(string(nextRune), nil)
			}
		}

		iter.Reset(startingIndex)
		return invalidMatchTree("", "SetOfCharacters:[" + characters + "], NoMatch: '" + string(nextRune) + "' not found in set")
	}
}
