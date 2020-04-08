package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConsecutiveCount(t *testing.T) {
	t.Run("if iter begins with a consecutive sequence of valid runes return the count", func(t *testing.T) {
		iter := CreateIterator("abcdefg875478574758")
		actual := 0
		require.Equal(t, 7, countConsecutiveCharacters(&iter, Let, &actual))
		require.Equal(t, '8', iter.Next())
	})

	t.Run("if iter begins with any invalid runes return 0", func(t *testing.T) {
		iter := CreateIterator("1abcdefg")
		actual := 0
		require.Equal(t, 0, countConsecutiveCharacters(&iter, Let, &actual))
		require.Equal(t, '1', iter.Next())
	})

	t.Run("if iter has only valid runes return the count", func(t *testing.T) {
		iter := CreateIterator("abcdefg")
		actual := 0
		require.Equal(t, 7, countConsecutiveCharacters(&iter, Let, &actual))
		require.Equal(t, false, iter.HasNext())
	})

	t.Run("if iter has no next then return 0", func(t *testing.T) {
		iter := CreateIterator("")
		actual := 0
		require.Equal(t, 0, countConsecutiveCharacters(&iter, Let, &actual))
	})
}
