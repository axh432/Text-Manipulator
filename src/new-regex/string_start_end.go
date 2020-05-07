package new_regex

func StringStart(iter *Iterator) MatchTree {
	return MatchTree{
		IsValid:  !iter.HasPrev(),
		Value:    "",
		Label:    "StringStart",
		Children: nil,
	}
}

func StringEnd(iter *Iterator) MatchTree {
	return MatchTree{
		IsValid:  !iter.HasNext(),
		Value:    "",
		Label:    "StringEnd",
		Children: nil,
	}
}
