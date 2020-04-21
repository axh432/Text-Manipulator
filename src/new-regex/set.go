package new_regex

func Set(expressions ...Expression) Expression {
	return func(iter *Iterator) bool {
		startingIndex := iter.GetIndex()
		for _, expression := range expressions {
			if expression(iter) {
				return true
			} else {
				iter.Reset(startingIndex)
			}
		}
		return false
	}
}
