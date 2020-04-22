package new_regex


import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("when given a string that matches an expression in a set return true", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)

		require.True(t, Match("a", exp).isValid)
		require.Equal(t, "a", Match("a", exp).Value)

		require.True(t, Match("b", exp).isValid)
		require.Equal(t, "b", Match("b", exp).Value)

		require.True(t, Match("c", exp).isValid)
		require.Equal(t, "c", Match("c", exp).Value)
	})

	t.Run("when given a character that is not in the set return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)
		require.False(t, Match("d", exp).isValid)
		require.Equal(t, "", Match("d", exp).Value)
	})

	t.Run("when given a string and the first character matches return true", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)
		require.True(t, Match("athguy", exp).isValid)
		require.Equal(t, "a", Match("athguy", exp).Value)
	})

	t.Run("when given a string and the first character does not match return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)
		require.False(t, Match("xthguy", exp).isValid)
		require.Equal(t, "", Match("xthguy", exp).Value)
	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)
		require.False(t, Match("", exp).isValid)
		require.Equal(t, "", Match("", exp).Value)
	})

	t.Run("when given an empty set return false", func(t *testing.T) {
		exp := Set()
		require.False(t, Match("a", exp).isValid)
		require.Equal(t, "", Match("a", exp).Value)
	})

}
