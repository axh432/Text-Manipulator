package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_match(t *testing.T) {
	iter := CreateIterator("voidvoid (l r i){+}")
	ovid := createCharacterSet("ovid")
	between4And8 := betweenNumberOfCharacters(4, 8)
	require.True(t, match(&iter, between4And8(ovid), One(White)))
}
