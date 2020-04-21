package new_regex

func Range(exp Expression, min int, max int) Expression {
	return func(iter *Iterator) bool {
		count := consecutiveCount(iter, exp)
		return isCountRightNumberOfCharacters(count, min, max)
	}
}

func consecutiveCount(iter *Iterator, exp Expression) int {
	count := 0
	for iter.HasNext() {
		if !exp(iter) {
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
