package new_regex

func Match(stringToMatch string, exp Expression) MatchTree {
	iter := CreateIterator(stringToMatch)
	return exp(&iter)
}
