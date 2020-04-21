package new_regex

func Not(exp Expression) Expression {
	return func(iter *Iterator) bool {
		return !exp(iter)
	}
}
