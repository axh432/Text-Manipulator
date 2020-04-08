package new_regex

func StringLiteral(str string) Expression {
	return func(iter *Iterator) bool {
		for _, char := range str {
			if !iter.HasNext() || char != iter.Next() {
				return false
			}
		}
		return true
	}
}
