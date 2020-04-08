package new_regex

import "unicode"

type CharacterSet func(r rune) bool

var (
	White = unicode.IsSpace
	Num   = unicode.IsNumber
	Let   = unicode.IsLetter
	Punct = unicode.IsPunct
	Sym   = unicode.IsSymbol
)
