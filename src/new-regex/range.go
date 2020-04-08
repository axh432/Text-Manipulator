package new_regex

func Range(set Expression, min int, max int) Expression {
	return func(iter *Iterator) bool {
		count := consecutiveCount(iter, set)
		return isCountRightNumberOfCharacters(count, min, max)
	}
}

func consecutiveCount(iter *Iterator, charSet Expression) int {
	count := 0
	for iter.HasNext() {
		if !charSet(iter) {
			iter.SafePrev()
			break
		}
		count++
	}
	return count
}

func isCountRightNumberOfCharacters(count, min, max int) bool {
	if max < 0 { //if unlimited
		return count >= min
	}
	return count >= min && count <= max
}
