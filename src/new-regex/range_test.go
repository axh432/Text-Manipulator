package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("when given a string that matches a range return true", func(t *testing.T) {
		a := SetOfCharacters("a")

		exp := Range(a, 1, 1)
		matchResult := Match2("a", exp)
		require.True(t, matchResult.isValid)
		require.Equal(t, "a", matchResult.Value)

		exp2 := Range(a, 1, 4)
		matchResult2 := Match2("aaa", exp2)
		require.True(t, matchResult2.isValid)
		require.Equal(t, "aaa", matchResult2.Value)

		exp3 := Range(a, 1, 4)
		matchResult3 := Match2("aaaa", exp3)
		require.True(t, matchResult3.isValid)
		require.Equal(t, "aaaa", matchResult3.Value)

		exp4 := Range(a, 1, 4)
		matchResult4 := Match2("a", exp4)
		require.True(t, matchResult4.isValid)
		require.Equal(t, "a", matchResult4.Value)
	})

	t.Run("when given a string that is greater than a range return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		exp := Range(a, 1, 1)
		matchResult := Match2("aa", exp)
		require.False(t, matchResult.isValid)
		require.Equal(t, "a", matchResult.Value)
	})

	t.Run("when given a string that is less than a range return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		exp := Range(a, 2, 2)
		require.False(t, Match2("a", exp).isValid)
		require.Equal(t, "a", Match2("a", exp).Value)
	})

	t.Run("when given a string that is empty", func(t *testing.T) {
		a := SetOfCharacters("a")
		exp := Range(a, 0, 1)
		require.True(t, Match2("", exp).isValid)

		exp2 := Range(a, 1, 1)
		require.False(t, Match2("", exp2).isValid)
	})

	t.Run("zero to many", func(t *testing.T) {
		a := SetOfCharacters("a")

		exp := Range(a, 0, -1)
		require.True(t, Match2("", exp).isValid)
		require.Equal(t, "", Match2("", exp).Value)

		require.True(t, Match2("aaaa", exp).isValid)
		require.Equal(t, "aaaa", Match2("aaaa", exp).Value)
	})

	t.Run("one to many", func(t *testing.T) {
		a := SetOfCharacters("a")

		exp := Range(a, 1, -1)
		require.False(t, Match2("", exp).isValid)
		require.Equal(t, "", Match2("", exp).Value)

		require.True(t, Match2("a", exp).isValid)
		require.Equal(t, "a", Match2("a", exp).Value)

		require.True(t, Match2("aaaaaaa", exp).isValid)
		require.Equal(t, "aaaaaaa", Match2("aaaaaaa", exp).Value)
	})
}
