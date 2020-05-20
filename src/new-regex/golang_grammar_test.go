package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_parse_go(t *testing.T) {

	t.Run("test function signature", func(t *testing.T) {
		require.True(t, Match("func copy()", functionSignature).IsValid)
		require.True(t, Match("func copy	()", functionSignature).IsValid)
		require.True(t, Match("func copy(	)", functionSignature).IsValid)
		require.True(t, Match("func copy()()", functionSignature).IsValid)
		require.True(t, Match("func copy(int left)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, \nint right)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right, float up)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) int", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (int, int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int, sedric int)", functionSignature).IsValid)
	})

	t.Run("strings", func(t *testing.T) {
		result := Match(`"fmt \" line\""`, String)
		require.True(t, result.IsValid)

	})

	t.Run("import", func(t *testing.T) {
		require.True(t, Match(`import "fmt"`, importDeclaration).IsValid)
		require.True(t, Match(`import ("fmt")`, importDeclaration).IsValid)
		require.True(t, Match("import (\"fmt\"\n\"strings\")", importDeclaration).IsValid)

	})

	t.Run("package", func(t *testing.T) {
		require.True(t, Match(`package somepackage`, packageDeclaration).IsValid)
	})

	t.Run("var assign statement", func(t *testing.T){
		require.True(t, Match(`var a = "initial"`, varAssignStatement).IsValid)
		require.True(t, Match(`var d = true"`, varAssignStatement).IsValid)
		require.True(t, Match(`var b int = 1`, varAssignStatement).IsValid)
		require.True(t, Match(`var b, c int = 1, 2"`, varAssignStatement).IsValid)
	})

	t.Run("assign statement", func(t *testing.T){
		require.True(t, Match(`f := "apple"`, assignStatement).IsValid)
	})

	t.Run("function call", func(t *testing.T){
		require.True(t, Match(`Println(a)`, functionCall).IsValid)
		require.True(t, Match(`fmt.Println(a)`, functionCall).IsValid)
		require.True(t, Match(`fmt.Println(b, c)`, functionCall).IsValid)
	})
}


