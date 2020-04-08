package new_regex

func StringToCharacterSet(characters string) Expression {
	return func(iter *Iterator) bool {
		if !iter.HasNext() {
			return false
		}
		nextRune := iter.Next()
		for _, char := range characters {
			if char == nextRune {
				return true
			}
		}
		return false
	}
}
