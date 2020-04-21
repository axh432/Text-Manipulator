package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetOfCharacters(t *testing.T) {
	t.Run("when given a character that is in the set return true", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.True(t, Match("a", exp))
		require.True(t, Match("b", exp))
		require.True(t, Match("c", exp))
	})

	t.Run("when given a character that is not in the set return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.False(t, Match("d", exp))
	})

	t.Run("when given a string and the first character matches return true", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.True(t, Match("athguy", exp))
	})

	t.Run("when given a string and the first character does not match return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.False(t, Match("xthguy", exp))
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.False(t, Match("", exp))
	})

	t.Run("when given an empty set return false", func(t *testing.T) {
		exp := SetOfCharacters("")
		require.False(t, Match("a", exp))
	})
}
