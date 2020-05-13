package new_regex

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_parse_go(t *testing.T) {

	t.Run("test function signature", func(t *testing.T) {
		require.True(t, Match("func copy()", functionSignature).IsValid)
		require.True(t, Match("func copy	()", functionSignature).IsValid)
		require.True(t, Match("func copy(	)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, \nint right)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right, float up)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) int", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (int, int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int, sedric int)", functionSignature).IsValid)
	})

	t.Run("lists", func(t *testing.T) {

		functionParametersList  := Sequence(openBracket, optionalWhitespaceNoNewLineBlock, parameterList, optionalWhitespaceNoNewLineBlock, closedBracket)

		result := Match("(int left, int right, float up)", functionParametersList)

		fmt.Println(result.toMermaidDiagram())

		require.True(t, result.IsValid)

	})
}


