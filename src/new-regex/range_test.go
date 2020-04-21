package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("when given a string that matches a range return true", func(t *testing.T) {
		a := SetOfCharacters("a")
		exp := Range(a, 1, 1)
		require.True(t, Match("a", exp))

		exp2 := Range(a, 1, 4)
		require.True(t, Match("aaa", exp2))

		exp3 := Range(a, 1, 4)
		require.True(t, Match("aaaa", exp3))

		exp4 := Range(a, 1, 4)
		require.True(t, Match("a", exp4))
	})

	t.Run("when given a string that is greater than a range return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		exp := Range(a, 1, 1)
		require.False(t, Match("aa", exp))
	})

	t.Run("when given a string that is less than a range return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		exp := Range(a, 2, 2)
		require.False(t, Match("a", exp))
	})

	t.Run("when given a string that is empty", func(t *testing.T) {
		a := SetOfCharacters("a")
		exp := Range(a, 0, 1)
		require.True(t, Match("", exp))

		exp2 := Range(a, 1, 1)
		require.False(t, Match("", exp2))
	})

	t.Run("zero to many", func(t *testing.T) {
		a := SetOfCharacters("a")

		exp := Range(a, 0, -1)
		require.True(t, Match("", exp))

		exp2 := Range(a, 0, -1)
		require.True(t, Match("aaaa", exp2))

	})

	t.Run("one to many", func(t *testing.T) {
		a := SetOfCharacters("a")

		exp := Range(a, 1, -1)
		require.False(t, Match("", exp))

		exp2 := Range(a, 1, -1)
		require.True(t, Match("a", exp2))

		exp3 := Range(a, 1, -1)
		require.True(t, Match("aaaaaaa", exp3))
	})
}
