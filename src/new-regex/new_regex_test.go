package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_match(t *testing.T) {
	stringToMatch := "voidvoid (l r i){+}"

	exp := Sequence(Range(StringToCharacterSet("ovid"), 4, 8), Range(Whitespace, 1, -1), StringLiteral("(l r i)"))

	require.True(t,
		Match(stringToMatch, exp))

}
