package new_regex

func createCharacterSet(characters string) CharacterSet {
	return func(r rune) bool {
		for _, char := range characters {
			if char == r {
				return true
			}
		}
		return false
	}
}

func combineSets(charSets ...CharacterSet) CharacterSet {
	return func(r rune) bool {
		for _, charSet := range charSets {
			if charSet(r) {
				return true
			}
		}
		return false
	}
}
