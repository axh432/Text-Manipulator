package new_regex

func StringStart(iter *Iterator) MatchTree {
	return MatchTree{
		isValid:  !iter.HasPrev(),
		Value:    "",
		Label:    "StringStart",
		Children: nil,
	}
}

func StringEnd(iter *Iterator) MatchTree {
	return MatchTree{
		isValid:  !iter.HasNext(),
		Value:    "",
		Label:    "StringEnd",
		Children: nil,
	}
}
