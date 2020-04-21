package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSequenceOfCharacters(t *testing.T) {
	t.Run("when given an exact string of characters this expression should return true", func(t *testing.T) {
		stringToMatch := "1234567890qwertyuiop[]asdfghjkl;'\\zxcvbnm`,./!@£$%^&*()_+{}|:?><~"

		exp := SequenceOfCharacters("1234567890qwertyuiop[]asdfghjkl;'\\zxcvbnm`,./!@£$%^&*()_+{}|:?><~")

		require.True(t,
			Match(stringToMatch, exp))
	})

	t.Run("when given a string of characters that differs in the beginning return false", func(t *testing.T) {
		stringToMatch := "a"

		exp := SequenceOfCharacters("b")

		require.False(t,
			Match(stringToMatch, exp))
	})

	t.Run("when given a string of characters that differs in the end return false", func(t *testing.T) {
		stringToMatch := "ab"

		exp := SequenceOfCharacters("ac")

		require.False(t,
			Match(stringToMatch, exp))
	})

	t.Run("when given a string of characters that differs in the middle return false", func(t *testing.T) {
		stringToMatch := "abc"

		exp := SequenceOfCharacters("adc")

		require.False(t,
			Match(stringToMatch, exp))
	})

	t.Run("when given a string of characters that is longer than the sequence return true", func(t *testing.T) {
		stringToMatch := "abcdefg"

		exp := SequenceOfCharacters("abc")

		require.True(t,
			Match(stringToMatch, exp))
	})

	t.Run("when given a string of characters that is shorter than the sequence return false", func(t *testing.T) {
		stringToMatch := "ab"

		exp := SequenceOfCharacters("abc")

		require.False(t,
			Match(stringToMatch, exp))
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		exp := SequenceOfCharacters("abc")
		require.False(t, Match("", exp))
	})

	t.Run("when given an empty sequence return false", func(t *testing.T) {
		exp := SequenceOfCharacters("")
		require.False(t, Match("a", exp))
	})
}