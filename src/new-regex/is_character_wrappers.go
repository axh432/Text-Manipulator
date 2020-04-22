package new_regex

import (
	"unicode"
)

type IsCharacterFunction func(r rune) bool

var (
	Whitespace  = createSetFromIsCharacterFunction(unicode.IsSpace)
	Number      = createSetFromIsCharacterFunction(unicode.IsNumber)
	Letter      = createSetFromIsCharacterFunction(unicode.IsLetter)
	Punctuation = createSetFromIsCharacterFunction(unicode.IsPunct)
	Symbol      = createSetFromIsCharacterFunction(unicode.IsSymbol)
)

func createSetFromIsCharacterFunction(isCharacterFunction IsCharacterFunction) Expression2 {
	return func(iter *Iterator) MatchTree {
		mt := MatchTree{}
		if iter.HasNext() {
			nextRune := iter.Next()
			if isCharacterFunction(nextRune) {
				mt.isValid = true
				mt.Value = string(nextRune)
			}
		}
		return mt
	}
}
