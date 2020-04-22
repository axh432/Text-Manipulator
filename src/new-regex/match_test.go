package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_match(t *testing.T) {

	t.Run("whitespace no new line", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")
		require.True(t, Match(" ", whitespaceNoNewLine).isValid)
		require.True(t, Match("\t", whitespaceNoNewLine).isValid)
		require.False(t, Match("\n", whitespaceNoNewLine).isValid)
		require.False(t, Match("", whitespaceNoNewLine).isValid)
		require.False(t, Match("a", whitespaceNoNewLine).isValid)
	})

	t.Run("optional whitespace no new line", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")
		optionalWhitespaceNoNewLine := Range(whitespaceNoNewLine, 0, 1)
		require.True(t, Match(" ", Sequence(optionalWhitespaceNoNewLine, StringEnd)).isValid)
		require.True(t, Match("\t", Sequence(optionalWhitespaceNoNewLine, StringEnd)).isValid)
		require.True(t, Match("", Sequence(optionalWhitespaceNoNewLine, StringEnd)).isValid)

		require.False(t, Match("\n", Sequence(optionalWhitespaceNoNewLine, StringEnd)).isValid)
		require.False(t, Match("a", Sequence(optionalWhitespaceNoNewLine, StringEnd)).isValid)
	})

	t.Run("java name", func(t *testing.T) {
		underscore := SetOfCharacters("_")
		name := Sequence(Set(Letter, underscore), Range(Set(Letter, Number, underscore), 0, -1))

		require.True(t, Match("m1", Sequence(name, StringEnd)).isValid)
		require.True(t, Match("myvariable", Sequence(name, StringEnd)).isValid)
		require.True(t, Match("myVariable", Sequence(name, StringEnd)).isValid)
		require.True(t, Match("my_variable", Sequence(name, StringEnd)).isValid)
		require.True(t, Match("my_variable_2", Sequence(name, StringEnd)).isValid)
		require.True(t, Match("_variable", Sequence(name, StringEnd)).isValid)

		require.False(t, Match("1variable", Sequence(name, StringEnd)).isValid)
		require.False(t, Match("my-variable-cool", Sequence(name, StringEnd)).isValid)
		require.Equal(t, "my", Match("my-variable-cool", Sequence(name, StringEnd)).Value)
	})

	t.Run("java function keywords", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")

		public := SequenceOfCharacters("public")
		private := SequenceOfCharacters("private")
		protected := SequenceOfCharacters("protected")
		static := Sequence(Range(whitespaceNoNewLine, 1, -1), SequenceOfCharacters("static"))

		functionKeywords := Sequence(Set(public, private, protected), Range(static, 0, 1))

		require.True(t, Match("public", functionKeywords).isValid)
		require.True(t, Match("private", functionKeywords).isValid)
		require.True(t, Match("protected", functionKeywords).isValid)
		require.True(t, Match("public static", functionKeywords).isValid)
		require.True(t, Match("private static", functionKeywords).isValid)
		require.True(t, Match("protected static", functionKeywords).isValid)

		require.False(t, Match("static public", functionKeywords).isValid)
		require.False(t, Match("something static", functionKeywords).isValid)
		require.False(t, Match("publicstatic", Sequence(functionKeywords, StringEnd)).isValid)
	})

	t.Run("java function parameters", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")
		underscore := SetOfCharacters("_")
		name := Sequence(Set(Letter, underscore), Range(Set(Letter, Number, underscore), 0, -1))
		typeName := name
		variableName := name

		parameter := Sequence(typeName, Range(whitespaceNoNewLine, 1, -1), variableName)

		require.True(t, Match("int count", parameter).isValid)
		require.True(t, Match("int 	count", parameter).isValid)
		require.True(t, Match("String name", parameter).isValid)

		require.False(t, Match("String", parameter).isValid)
		require.False(t, Match("String ", parameter).isValid)
	})

	/*t.Run("java list of function parameters", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")
		underscore := SetOfCharacters("_")
		comma := SetOfCharacters(",")
		name := Sequence(Set(Letter, underscore), Range(Set(Letter, Number, underscore), 0, -1))
		typeName := name
		variableName := name
		parameter := Sequence(typeName, Range(whitespaceNoNewLine, 1, -1), variableName)

		parameterComma := Sequence(parameter, Range(whitespaceNoNewLine, 0, -1), comma)

		listOfParameters := Sequence(parameterComma, Range(Sequence(whitespaceNoNewLine, parameterComma), 0, -1), parameter)

		result := Match("int width, int length", listOfParameters)

		fmt.Println(result.toString())

		require.True(t, Match("int width, int length", listOfParameters).isValid)
		require.True(t, Match("int width  , int length", listOfParameters).isValid)

	})*/

	/*t.Run("sophisticated match the string", func(t *testing.T) {

		stringToMatch := "public void myFunction(int left, int right) int"

		//Primitives
		underscore := SetOfCharacters("_")
		comma := SetOfCharacters(",")
		openbracket := SetOfCharacters("(")
		closedBracket := SetOfCharacters(")")

		whitespaceNoNewLine := SetOfCharacters(" \t")
		optionalWhitespaceNoNewLine := Range(whitespaceNoNewLine, 0, -1)

		name := Sequence(Letter, Range(Set(Letter, Number, underscore), 0, -1))

		//Java
		publicKeyword := SequenceOfCharacters("public")
		voidKeyword := SequenceOfCharacters("void")
		keywords := Sequence(publicKeyword, whitespaceNoNewLine, voidKeyword)

		functionName := name
		typeName := name
		variableName := name
		returnType := name

		parameter := Sequence(typeName, whitespaceNoNewLine, variableName)
		parameterComma := Sequence(parameter, optionalWhitespaceNoNewLine, comma)

		listOfParameters := Sequence(parameterComma, Range(Sequence(whitespaceNoNewLine, parameterComma), 0, -1), parameter)
		singleParameter := parameter
		noParameters := optionalWhitespaceNoNewLine

		functionParameters := Sequence(openbracket, optionalWhitespaceNoNewLine, Set(listOfParameters, singleParameter, noParameters), optionalWhitespaceNoNewLine, closedBracket)

		functionSignature := Sequence(keywords, whitespaceNoNewLine, functionName, functionParameters, optionalWhitespaceNoNewLine, returnType)

		result := Match(stringToMatch, functionSignature)

		println(result.toString())
		require.True(t, result.isValid)
	})*/

}
