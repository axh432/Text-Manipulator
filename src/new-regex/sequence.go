package new_regex

func Sequence(expressions ...Expression) Expression {
	return func(iter *Iterator) bool {

		if len(expressions) == 0 {
			return false
		}

		for _, exp := range expressions {
			if !exp(iter) {
				return false
			}
		}
		return true
	}
}
