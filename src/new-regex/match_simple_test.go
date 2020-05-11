package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_match_simple(t *testing.T) {

	t.Run("simple match string", func(t *testing.T) {
		stringToMatch := "void myFunc(l r) { return l + r }"

		void := SequenceOfCharacters("void")
		funcName := SequenceOfCharacters("myFunc")
		params := SequenceOfCharacters("(l r)")
		body := SequenceOfCharacters("{ return l + r }")

		exp := Sequence(void, Whitespace, funcName, params, Whitespace, body)

		result := Match(stringToMatch, exp)

		println(result.toMermaidDiagram())
		require.True(t, result.IsValid)
		require.Equal(t, stringToMatch, result.Value)

	})

	t.Run("simple label nodes of interest", func(t *testing.T) {
		stringToMatch := "void myFunc(l r) { return l + r }"

		void := Label(SequenceOfCharacters("void"), "void")
		funcName := Label(SequenceOfCharacters("myFunc"), "funcName")
		params := Label(SequenceOfCharacters("(l r)"), "params")
		body := Label(SequenceOfCharacters("{ return l + r }"), "body")

		exp := Sequence(void, Whitespace, funcName, params, Whitespace, body)

		result := Match(stringToMatch, exp)

		println(result.toString())
		require.True(t, result.IsValid)
		require.Equal(t, stringToMatch, result.Value)
	})
}
