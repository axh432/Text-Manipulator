package new_regex

type Expression func(iter *Iterator) bool

func match(iter *Iterator, expressions ...Expression) bool {
	result := true
	for _, exp := range expressions {
		result = result && exp(iter)
	}
	return result
}
