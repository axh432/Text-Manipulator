package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSequenceOfCharacters(t *testing.T) {
	t.Run("when given an exact string of characters this expression should return true", func(t *testing.T) {
		stringToMatch := "1234567890qwertyuiop[]asdfghjkl;'\\zxcvbnm`,./!@£$%^&*()_+{}|:?><~"
		exp := SequenceOfCharacters("1234567890qwertyuiop[]asdfghjkl;'\\zxcvbnm`,./!@£$%^&*()_+{}|:?><~")
		matchResult := Match2(stringToMatch, exp)
		require.True(t, matchResult.isValid)
		require.Equal(t, stringToMatch, matchResult.Value)
	})

	t.Run("when given a string of characters that differs in the beginning return false", func(t *testing.T) {
		stringToMatch := "a"
		exp := SequenceOfCharacters("b")
		matchResult := Match2(stringToMatch, exp)
		require.False(t, matchResult.isValid)
		require.Equal(t, "", matchResult.Value)
	})

	t.Run("when given a string of characters that differs in the end return false", func(t *testing.T) {
		stringToMatch := "ab"
		exp := SequenceOfCharacters("ac")
		matchResult := Match2(stringToMatch, exp)
		require.False(t, matchResult.isValid)
		require.Equal(t, "a", matchResult.Value)
	})

	t.Run("when given a string of characters that differs in the middle return false", func(t *testing.T) {
		stringToMatch := "abc"
		exp := SequenceOfCharacters("adc")
		matchResult := Match2(stringToMatch, exp)
		require.False(t, matchResult.isValid)
		require.Equal(t, "a", matchResult.Value)
	})

	t.Run("when given a string of characters that is longer than the sequence return true", func(t *testing.T) {
		stringToMatch := "abcdefg"
		exp := SequenceOfCharacters("abc")
		matchResult := Match2(stringToMatch, exp)
		require.True(t, matchResult.isValid)
		require.Equal(t, "abc", matchResult.Value)
	})

	t.Run("when given a string of characters that is shorter than the sequence return false", func(t *testing.T) {
		stringToMatch := "ab"
		exp := SequenceOfCharacters("abc")
		matchResult := Match2(stringToMatch, exp)
		require.False(t, matchResult.isValid)
		require.Equal(t, "ab", matchResult.Value)
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		exp := SequenceOfCharacters("abc")
		matchResult := Match2("", exp)
		require.False(t, matchResult.isValid)
		require.Equal(t, "", matchResult.Value)
	})

	t.Run("when given an empty sequence return false", func(t *testing.T) {
		exp := SequenceOfCharacters("")
		matchResult := Match2("a", exp)
		require.False(t, matchResult.isValid)
		require.Equal(t, "", matchResult.Value)
	})
}
