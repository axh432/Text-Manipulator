package new_regex

type NumberOfCharacters func(charSet CharacterSet) Expression

var (
	ZeroOrMore = specifyNumberOfCharacters(0, GreaterThanOrEqual)
	OneOrMore  = specifyNumberOfCharacters(1, GreaterThanOrEqual)
	One        = specifyNumberOfCharacters(1, Equal)
)

func betweenNumberOfCharacters(min, max int) NumberOfCharacters {
	return func(charSet CharacterSet) Expression {
		return func(iter *Iterator) bool {
			count := 0
			countConsecutiveCharacters(iter, charSet, &count)
			return count >= min && count <= max
		}
	}
}

func specifyNumberOfCharacters(numOfChars int, operator EqualityOperator) NumberOfCharacters {
	return func(charSet CharacterSet) Expression {
		return func(iter *Iterator) bool {
			count := 0
			countConsecutiveCharacters(iter, charSet, &count)
			return operator(count, numOfChars)
		}
	}
}
