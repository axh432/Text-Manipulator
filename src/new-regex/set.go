package new_regex

func Set(expressions ...Expression) Expression {
	return func(iter *Iterator) bool {
		for _, expression := range expressions {
			if expression(iter) {
				return true
			}
		}
		return false
	}
}
