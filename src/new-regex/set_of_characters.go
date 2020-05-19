package new_regex

import (
	"strings"
	"unicode"
)

func SetOfCharacters(characters string) Expression {
	return func(iter *Iterator) MatchTree {
		if !iter.HasNext() {
			return invalidMatchTree("", "SetOfCharacters", nil, "SetOfCharacters:["+characters+"], NoMatch:reached end of string before finished")
		}

		startingIndex := iter.index
		nextRune := iter.Next()
		for _, char := range characters {
			if char == nextRune {
				return validMatchTree(string(nextRune), "SetOfCharacters", nil)
			}
		}

		iter.Reset(startingIndex)
		return invalidMatchTree("", "SetOfCharacters", nil, "SetOfCharacters:["+characters+"], NoMatch: '"+string(nextRune)+"' not found in set")
	}
}

func SetOfNotCharacters(characters string) Expression {
	return func(iter *Iterator) MatchTree {
		if !iter.HasNext() {
			return invalidMatchTree("", "SetOfNotCharacters", nil, "SetOfNotCharacters:["+characters+"], NoMatch:reached end of string before finished")
		}

		startingIndex := iter.index
		nextRune := iter.Next()
		for _, char := range characters {
			if char == nextRune {
				iter.Reset(startingIndex)
				return invalidMatchTree("", "SetOfNotCharacters", nil, "SetOfNotCharacters:["+characters+"], NoMatch: '"+string(nextRune)+"' found in set")
			}
		}

		return validMatchTree(string(nextRune), "SetOfNotCharacters", nil)
	}
}

func getRangeOfLettersLower(from rune, to rune) string {

	if !unicode.IsLetter(from) || !unicode.IsLetter(to) {
		return ""
	}

	from = unicode.ToLower(from)
	to = unicode.ToLower(to)

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	start := strings.IndexRune(alphabet, from)
	end   := strings.IndexRune(alphabet, to)
	return alphabet[start:end+1]
}

func getRangeOfDigits(from rune, to rune) string {

	if !unicode.IsDigit(from) || !unicode.IsDigit(to) {
		return ""
	}

	alphabet := "0123456789"
	start := strings.IndexRune(alphabet, from)
	end   := strings.IndexRune(alphabet, to)
	return alphabet[start:end+1]
}

func getRangeOfLettersUpper(from rune, to rune) string {

	if !unicode.IsLetter(from) || !unicode.IsLetter(to) {
		return ""
	}

	from = unicode.ToUpper(from)
	to = unicode.ToUpper(to)

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	start := strings.IndexRune(alphabet, from)
	end   := strings.IndexRune(alphabet, to)
	return alphabet[start:end+1]
}
