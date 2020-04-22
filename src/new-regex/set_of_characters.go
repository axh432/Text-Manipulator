package new_regex

func SetOfCharacters(characters string) Expression {
	return func(iter *Iterator) MatchTree {
		mt := MatchTree{}
		if !iter.HasNext() {
			mt.DebugLine = "SetOfCharacters:[" + characters + "], NoMatch:reached end of string before finished"
			return mt
		}
		nextRune := iter.Next()
		for _, char := range characters {
			if char == nextRune {
				mt.isValid = true
				mt.Value = string(nextRune)
				return mt
			}
		}
		mt.DebugLine = "SetOfCharacters:[" + characters + "], NoMatch: '" + string(nextRune) + "' not found in set"
		return mt
	}
}
