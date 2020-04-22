package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetOfCharacters(t *testing.T) {
	t.Run("when given a character that is in the set return true", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.True(t, Match2("a", exp).isValid)
		require.Equal(t, "a", Match2("a", exp).Value)

		require.True(t, Match2("b", exp).isValid)
		require.Equal(t, "b", Match2("b", exp).Value)

		require.True(t, Match2("c", exp).isValid)
		require.Equal(t, "c", Match2("c", exp).Value)
	})

	t.Run("when given a character that is not in the set return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.False(t, Match2("d", exp).isValid)
		require.Equal(t, "", Match2("d", exp).Value)
	})

	t.Run("when given a string and the first character matches return true", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.True(t, Match2("athguy", exp).isValid)
		require.Equal(t, "a", Match2("athguy", exp).Value)
	})

	t.Run("when given a string and the first character does not match return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.False(t, Match2("xthguy", exp).isValid)
		require.Equal(t, "", Match2("xthguy", exp).Value)
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		exp := SetOfCharacters("abc")
		require.False(t, Match2("", exp).isValid)
		require.Equal(t, "", Match2("", exp).Value)
	})

	t.Run("when given an empty set return false", func(t *testing.T) {
		exp := SetOfCharacters("")
		require.False(t, Match2("a", exp).isValid)
		require.Equal(t, "", Match2("a", exp).Value)
	})
}
