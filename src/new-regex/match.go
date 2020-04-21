package new_regex

type Expression func(iter *Iterator) bool

func Match(stringToMatch string, expression Expression) bool {
	iter := CreateIterator(stringToMatch)
	return expression(&iter)
}
