package new_regex

func Match(stringToMatch string, expression Expression) bool {
	iter := CreateIterator(stringToMatch)
	return expression(&iter)
}

func Match2(stringToMatch string, expression Expression2) MatchTree {
	iter := CreateIterator(stringToMatch)
	return expression(&iter)
}
