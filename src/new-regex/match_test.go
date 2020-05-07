package new_regex

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_match(t *testing.T) {

	t.Run("whitespace no new line", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")
		require.True(t, Match(" ", whitespaceNoNewLine).IsValid)
		require.True(t, Match("\t", whitespaceNoNewLine).IsValid)
		require.False(t, Match("\n", whitespaceNoNewLine).IsValid)
		require.False(t, Match("", whitespaceNoNewLine).IsValid)
		require.False(t, Match("a", whitespaceNoNewLine).IsValid)
	})

	t.Run("optional whitespace no new line", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")
		optionalWhitespaceNoNewLine := Range(whitespaceNoNewLine, 0, 1)
		require.True(t, Match(" ", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)
		require.True(t, Match("\t", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)
		require.True(t, Match("", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)

		require.False(t, Match("\n", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)
		require.False(t, Match("a", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)
	})

	t.Run("java name", func(t *testing.T) {
		underscore := SetOfCharacters("_")
		name := Sequence(Set(Letter, underscore), Range(Set(Letter, Number, underscore), 0, -1))

		require.True(t, Match("m1", Sequence(name, StringEnd)).IsValid)
		require.True(t, Match("myvariable", Sequence(name, StringEnd)).IsValid)
		require.True(t, Match("myVariable", Sequence(name, StringEnd)).IsValid)
		require.True(t, Match("my_variable", Sequence(name, StringEnd)).IsValid)
		require.True(t, Match("my_variable_2", Sequence(name, StringEnd)).IsValid)
		require.True(t, Match("_variable", Sequence(name, StringEnd)).IsValid)

		require.False(t, Match("1variable", Sequence(name, StringEnd)).IsValid)
		require.False(t, Match("my-variable-cool", Sequence(name, StringEnd)).IsValid)
		require.Equal(t, "my", Match("my-variable-cool", Sequence(name, StringEnd)).Value)
	})

	t.Run("java function keywords", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")

		public := SequenceOfCharacters("public")
		private := SequenceOfCharacters("private")
		protected := SequenceOfCharacters("protected")
		static := Sequence(Range(whitespaceNoNewLine, 1, -1), SequenceOfCharacters("static"))

		functionKeywords := Sequence(Set(public, private, protected), Range(static, 0, 1))

		require.True(t, Match("public", functionKeywords).IsValid)
		require.True(t, Match("private", functionKeywords).IsValid)
		require.True(t, Match("protected", functionKeywords).IsValid)
		require.True(t, Match("public static", functionKeywords).IsValid)
		require.True(t, Match("private static", functionKeywords).IsValid)
		require.True(t, Match("protected static", functionKeywords).IsValid)

		require.False(t, Match("static public", functionKeywords).IsValid)
		require.False(t, Match("something static", functionKeywords).IsValid)
		require.False(t, Match("publicstatic", Sequence(functionKeywords, StringEnd)).IsValid)
	})

	t.Run("java function parameters", func(t *testing.T) {
		whitespaceNoNewLine := SetOfCharacters(" \t")
		underscore := SetOfCharacters("_")
		name := Sequence(Set(Letter, underscore), Range(Set(Letter, Number, underscore), 0, -1))
		typeName := name
		variableName := name

		parameter := Sequence(typeName, Range(whitespaceNoNewLine, 1, -1), variableName)

		require.True(t, Match("int count", parameter).IsValid)
		require.True(t, Match("int 	count", parameter).IsValid)
		require.True(t, Match("String name", parameter).IsValid)

		require.False(t, Match("String", parameter).IsValid)
		require.False(t, Match("String ", parameter).IsValid)
	})

	t.Run("java list of function parameters", func(t *testing.T) {
		whitespaceNoNewLine := Label(SetOfCharacters(" \t"), "whitespaceNoNewLine")
		underscore := SetOfCharacters("_")
		comma := SetOfCharacters(",")
		name := Sequence(Set(Letter, underscore), Range(Set(Letter, Number, underscore), 0, -1))
		typeName := Label(name, "typename")
		variableName := Label(name, "variableName")
		parameter := Label(Sequence(typeName, Range(whitespaceNoNewLine, 1, -1), variableName), "parameter")

		parameterComma := Label(Sequence(parameter, Range(whitespaceNoNewLine, 0, -1), comma), "parameterComma")

		parameterCommaRange := Label(Range(Sequence(whitespaceNoNewLine, parameterComma), 0, -1), "parameterCommaRange")

		listOfParameters := Sequence(parameterComma, parameterCommaRange, Range(whitespaceNoNewLine, 0, -1), parameter)

		//listOfParameters := Sequence(parameterComma, parameterCommaRange, parameter)

		result := Match("int width, int length", listOfParameters)

		fmt.Println(result.DebugLine)

		require.True(t, Match("int width, int length", listOfParameters).IsValid)
		require.True(t, Match("int width, int length, string name", listOfParameters).IsValid)
		//require.True(t, Match("int width  , int length", listOfParameters).IsValid)

	})

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
		require.True(t, result.IsValid)
	})*/

}
