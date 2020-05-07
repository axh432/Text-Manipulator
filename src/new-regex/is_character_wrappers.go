package new_regex

import (
	"unicode"
)

//Todo: needs to be refactored in the light of 'type' and 'debugline' e.c.t

type IsCharacterFunction func(r rune) bool

var (
	Whitespace  = createSetFromIsCharacterFunction(unicode.IsSpace)
	Number      = createSetFromIsCharacterFunction(unicode.IsNumber)
	Letter      = createSetFromIsCharacterFunction(unicode.IsLetter)
	Punctuation = createSetFromIsCharacterFunction(unicode.IsPunct)
	Symbol      = createSetFromIsCharacterFunction(unicode.IsSymbol)
)

func createSetFromIsCharacterFunction(isCharacterFunction IsCharacterFunction) Expression {
	return func(iter *Iterator) MatchTree {
		mt := MatchTree{}
		if iter.HasNext() {
			nextRune := iter.Next()
			if isCharacterFunction(nextRune) {
				mt.IsValid = true
				mt.Value = string(nextRune)
			}
		}
		return mt
	}
}
