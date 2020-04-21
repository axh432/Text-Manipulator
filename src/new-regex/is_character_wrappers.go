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

func createSetFromIsCharacterFunction(isCharacterFunction IsCharacterFunction) Expression {
	return func(iter *Iterator) bool {
		if iter.HasNext() {
			return isCharacterFunction(iter.Next())
		}
		return false
	}
}
