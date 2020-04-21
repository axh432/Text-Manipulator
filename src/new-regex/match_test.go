package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_match(t *testing.T) {

	t.Run("match the string", func(t *testing.T) {
		stringToMatch := "voidvoid (l r i){+}"

		exp := Sequence(Range(SetOfCharacters("ovid"), 4, 8), Whitespace, SequenceOfCharacters("(l r i)"))

		require.True(t,
			Match(stringToMatch, exp))
	})

}
