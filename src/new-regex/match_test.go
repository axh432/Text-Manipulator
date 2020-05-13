package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_match(t *testing.T) {

	//Primitives
	underscore := SetOfCharacters("_")
	comma := SetOfCharacters(",")
	openbracket := SetOfCharacters("(")
	closedBracket := SetOfCharacters(")")
	whitespaceNoNewLine := SetOfCharacters(" \t")
	optionalWhitespaceNoNewLine := Range(whitespaceNoNewLine, 0, -1)
	//Java
	name := Sequence(Set(Letter, underscore), Range(Set(Letter, Number, underscore), 0, -1))
	public := SequenceOfCharacters("public")
	private := SequenceOfCharacters("private")
	protected := SequenceOfCharacters("protected")
	static := SequenceOfCharacters("static")
	functionKeywords := Sequence(Set(public, private, protected), Range(Sequence(whitespaceNoNewLine, static), 0, 1))
	typeName := name
	variableName := name
	returnType := name
	functionName := name
	parameter := Sequence(typeName, Range(whitespaceNoNewLine, 1, -1), variableName)
	parameterComma := Label(Sequence(parameter, optionalWhitespaceNoNewLine, comma), "parameterComma")
	parameterCommaRange := Label(Range(Sequence(whitespaceNoNewLine, parameterComma), 0, -1), "parameterCommaRange")
	noParameters := optionalWhitespaceNoNewLine
	listOfParameters := Sequence(parameterCommaRange, optionalWhitespaceNoNewLine, parameter)
	functionParameters := Sequence(openbracket, optionalWhitespaceNoNewLine, Set(listOfParameters, parameter, noParameters), optionalWhitespaceNoNewLine, closedBracket)
	functionSignature := Sequence(functionKeywords, whitespaceNoNewLine, returnType, whitespaceNoNewLine, functionName, functionParameters)

	t.Run("whitespace no new line", func(t *testing.T) {
		require.True(t, Match(" ", whitespaceNoNewLine).IsValid)
		require.True(t, Match("\t", whitespaceNoNewLine).IsValid)
		require.False(t, Match("\n", whitespaceNoNewLine).IsValid)
		require.False(t, Match("", whitespaceNoNewLine).IsValid)
		require.False(t, Match("a", whitespaceNoNewLine).IsValid)
	})

	t.Run("optional whitespace no new line", func(t *testing.T) {
		require.True(t, Match(" ", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)
		require.True(t, Match("\t", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)
		require.True(t, Match("", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)

		require.False(t, Match("\n", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)
		require.False(t, Match("a", Sequence(optionalWhitespaceNoNewLine, StringEnd)).IsValid)
	})

	t.Run("java name", func(t *testing.T) {
		require.True(t, Match("m1", name).IsValid)
		require.True(t, Match("myvariable", name).IsValid)
		require.True(t, Match("myVariable", name).IsValid)
		require.True(t, Match("my_variable", name).IsValid)
		require.True(t, Match("my_variable_2", name).IsValid)
		require.True(t, Match("_variable", name).IsValid)

		require.False(t, Match("1variable", Sequence(name, StringEnd)).IsValid)
		require.False(t, Match("my-cool-variable", Sequence(name, StringEnd)).IsValid)
	})

	t.Run("java function keywords", func(t *testing.T) {
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
		require.True(t, Match("int count", parameter).IsValid)
		require.True(t, Match("int 	count", parameter).IsValid)
		require.True(t, Match("String name", parameter).IsValid)

		require.False(t, Match("String", parameter).IsValid)
		require.False(t, Match("String ", parameter).IsValid)
	})

	t.Run("java list of function parameters", func(t *testing.T) {
		require.True(t, Match("int width, int length", listOfParameters).IsValid)
		require.True(t, Match("int width, int length, string name", listOfParameters).IsValid)
		require.True(t, Match("int width  , int length", listOfParameters).IsValid)
		require.True(t, Match("int length", listOfParameters).IsValid)
	})

	t.Run("sophisticated match the string", func(t *testing.T) {

		//Primitives

		//Java
		result := Match("public void myFunction(int left)", functionSignature)
		println(result.toMermaidDiagram())

		require.True(t, Match("public void myFunction(int left, int right)", functionSignature).IsValid)
	})

}
