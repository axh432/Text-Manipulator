package new_regex

import (
	. "text-manipulator/src/functional"
)

func countConsecutiveCharacters(iter *Iterator, charSet CharacterSet, count *int) Statement {
	return Whilst(iter.HasNext,
		If(nextCharacterIsNotInSet(iter, charSet),
			moveIteratorBackAndReturn(iter),
			Else(Increment(count))))
}

func nextCharacterIsNotInSet(iter *Iterator, inCharSet CharacterSet) Statement {
	return func() bool {
		return !inCharSet(iter.Next())
	}
}

func moveIteratorBackAndReturn(iter *Iterator) Statement {
	return func() bool {
		iter.SafePrev()
		return true
	}
}

func Increment(num *int) Statement {
	return func() bool {
		*num++
		return false
	}
}
