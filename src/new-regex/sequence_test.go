package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSequence(t *testing.T) {
	t.Run("when given an exact string of characters this expression should return true", func(t *testing.T) {

		stringToMatch := "abc"

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)

		matchResult := Match(stringToMatch, exp)

		require.True(t, matchResult.isValid)
		require.Equal(t, stringToMatch, matchResult.Value)
	})

	t.Run("when given a string of characters that differs in the beginning return false", func(t *testing.T) {
		stringToMatch := "bbc"

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)

		matchResult := Match(stringToMatch, exp)

		require.False(t, matchResult.isValid)
		require.Equal(t, "", matchResult.Value)
	})

	t.Run("when given a string of characters that differs in the end return false", func(t *testing.T) {
		stringToMatch := "abb"

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)

		matchResult := Match(stringToMatch, exp)

		require.False(t, matchResult.isValid)
		require.Equal(t, "ab", matchResult.Value)
	})

	t.Run("when given a string of characters that differs in the middle return false", func(t *testing.T) {
		stringToMatch := "aac"

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)

		matchResult := Match(stringToMatch, exp)

		require.False(t, matchResult.isValid)
		require.Equal(t, "a", matchResult.Value)
	})

	t.Run("when given a string of characters that is longer than the sequence return true", func(t *testing.T) {
		stringToMatch := "abcdefg"

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)

		matchResult := Match(stringToMatch, exp)

		require.True(t, matchResult.isValid)
		require.Equal(t, "abc", matchResult.Value)
	})

	t.Run("when given a string of characters that is shorter than the sequence return false", func(t *testing.T) {
		stringToMatch := "ab"

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)

		matchResult := Match(stringToMatch, exp)

		require.False(t, matchResult.isValid)
		require.Equal(t, "ab", matchResult.Value)
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Sequence(a, b, c)

		matchResult := Match("", exp)

		require.False(t, matchResult.isValid)
		require.Equal(t, "", matchResult.Value)
	})

	t.Run("when given an empty sequence return false", func(t *testing.T) {
		exp := Sequence()
		matchResult := Match("a", exp)

		require.False(t, matchResult.isValid)
		require.Equal(t, "", matchResult.Value)
	})
}
