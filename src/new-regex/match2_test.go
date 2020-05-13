package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//either A or AB
func firstOrBoth(A Expression, B Expression) Expression {
	return Set(A, Sequence(A,B))
}

//covers the patterns ABA, ABABA e.c.t
func repeatingList(A Expression, B Expression) Expression {
	return Sequence(Range(Sequence(A, B), 1, -1), A)
}

func blockOf(A Expression) Expression {
	return Range(A, 1, -1)
}

func Test_match2(t *testing.T) {

	//Primitives
	whitespaceNoNewLine := blockOf(SetOfCharacters(" \t"))
	underscore := SetOfCharacters("_")
	comma := SetOfCharacters(",")
	commaSpace := firstOrBoth(comma, whitespaceNoNewLine)
	openBracket := SetOfCharacters("(")
	openBracketSpace := firstOrBoth(openBracket, whitespaceNoNewLine)
	closedBracket := SetOfCharacters(")")
	closedBracketSpace := firstOrBoth(closedBracket, whitespaceNoNewLine)

	//Java
	blockOfLetterNumberUnderscore := blockOf(Set(Letter, Number, underscore))
	letterOrUnderscore := Set(Letter, underscore)
	name := firstOrBoth(letterOrUnderscore, blockOfLetterNumberUnderscore)

	typeName := Label(name, "typename")
	variableName := Label(name, "variablename")
	returnType := Label(name, "returntype")
	functionName := Label(name, "functionName")

	public := SequenceOfCharacters("public")
	private := SequenceOfCharacters("private")
	protected := SequenceOfCharacters("protected")
	static := SequenceOfCharacters("static")

	accessSpecifier := Sequence(Set(public, private, protected), whitespaceNoNewLine)
	staticSpace := Sequence(static, whitespaceNoNewLine)
	functionKeywords := firstOrBoth(accessSpecifier, staticSpace)

	parameter := Label(Sequence(typeName, whitespaceNoNewLine, variableName), "parameter")
	parameterSpace := firstOrBoth(parameter, whitespaceNoNewLine)
	parameterList := repeatingList(parameterSpace, commaSpace)

	functionParametersList := Sequence(openBracketSpace, parameterList, closedBracketSpace)
	functionParametersSingle := Sequence(openBracketSpace, parameterSpace, closedBracketSpace)
	functionParametersEmpty := Sequence(openBracketSpace, closedBracketSpace)

	functionParameters := Set(functionParametersList, functionParametersSingle, functionParametersEmpty)

	functionSignature := Sequence(functionKeywords, returnType, whitespaceNoNewLine, functionName, functionParameters)

	t.Run("whitespace no new line", func(t *testing.T) {
		require.True(t, Match(" ", whitespaceNoNewLine).IsValid)
		require.True(t, Match("\t", whitespaceNoNewLine).IsValid)
		require.False(t, Match("\n", whitespaceNoNewLine).IsValid)
		require.False(t, Match("", whitespaceNoNewLine).IsValid)
		require.False(t, Match("a", whitespaceNoNewLine).IsValid)
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
		require.True(t, Match("public ", functionKeywords).IsValid)
		require.True(t, Match("private ", functionKeywords).IsValid)
		require.True(t, Match("protected ", functionKeywords).IsValid)
		require.True(t, Match("public static ", functionKeywords).IsValid)
		require.True(t, Match("private static ", functionKeywords).IsValid)
		require.True(t, Match("protected static ", functionKeywords).IsValid)

		require.False(t, Match("static public ", functionKeywords).IsValid)
		require.False(t, Match("something static ", functionKeywords).IsValid)
		require.False(t, Match("publicstatic ", Sequence(functionKeywords, StringEnd)).IsValid)
	})

	t.Run("java function parameters", func(t *testing.T) {
		require.True(t, Match("int count", parameter).IsValid)
		require.True(t, Match("int 	count", parameter).IsValid)
		require.True(t, Match("String name", parameter).IsValid)

		require.False(t, Match("String", parameter).IsValid)
		require.False(t, Match("String ", parameter).IsValid)
	})

	t.Run("java list of function parameters", func(t *testing.T) {
		require.True(t, Match("int width, int length", parameterList).IsValid)
		require.True(t, Match("int width, int length, string name", parameterList).IsValid)
		require.True(t, Match("int width  , int length", parameterList).IsValid)
		require.False(t, Match("int length", parameterList).IsValid)
	})

	t.Run("sophisticated match the string", func(t *testing.T) {
		require.True(t, Match("public void delete(int left, int right)", functionSignature).IsValid)
		require.True(t, Match("public void delete(int left , int right)", functionSignature).IsValid)
		require.True(t, Match("public void delete( int left, int right)", functionSignature).IsValid)
		require.True(t, Match("public void delete(int left, int right )", functionSignature).IsValid)
		require.True(t, Match("public void delete()", functionSignature).IsValid)
		require.True(t, Match("public void delete( )", functionSignature).IsValid)
		require.True(t, Match("public void delete(int left)", functionSignature).IsValid)
		require.True(t, Match("public static void copy(int left)", functionSignature).IsValid)
		require.True(t, Match("private void copy(int left)", functionSignature).IsValid)
		require.True(t, Match("private static void copy(int left)", functionSignature).IsValid)
		require.True(t, Match("protected void copy(int left)", functionSignature).IsValid)
		require.True(t, Match("protected static void copy(int left)", functionSignature).IsValid)
	})

}
