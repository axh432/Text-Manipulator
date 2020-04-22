package new_regex

func StringStart(iter *Iterator) bool {
	return !iter.HasPrev()
}

func StringEnd(iter *Iterator) bool {
	return !iter.HasNext()
}
