package new_regex

func StringStart(iter *Iterator) MatchTree {
	return MatchTree{
		isValid:  !iter.HasPrev(),
		Value:    "",
		Name:     "StringStart",
		Children: nil,
	}
}

func StringEnd(iter *Iterator) MatchTree {
	return MatchTree{
		isValid:  !iter.HasNext(),
		Value:    "",
		Name:     "StringEnd",
		Children: nil,
	}
}
