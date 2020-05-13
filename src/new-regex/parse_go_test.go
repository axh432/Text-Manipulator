package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//either A or AB
/*func firstOrBoth(A Expression, B Expression) Expression {
	return Set(A, Sequence(A,B))
}

//covers the patterns ABA, ABABA e.c.t
func repeatingList(A Expression, B Expression) Expression {
	return Sequence(Range(Sequence(A, B), 1, -1), A)
}

func blockOf(A Expression) Expression {
	return Range(A, 1, -1)
}*/

func optionalWhitespaceNoNewLine(A Expression) Expression {
	return firstOrBoth(A, Range(SetOfCharacters(" \t"), 1, -1))
}

func Test_parse_go(t *testing.T) {

	//Primitives
	whitespaceNoNewLine := Range(SetOfCharacters(" \t"), 1, -1)
	//newline := SequenceOfCharacters("\n")
	//whitespaceAtLeastOneNewLine := Sequence(whitespaceNoNewLine, newline, Whitespace)
	underscore := SetOfCharacters("_")
	comma := SetOfCharacters(",")
	commaOptionalSpace := optionalWhitespaceNoNewLine(comma)
	openBracket := SetOfCharacters("(")
	openBracketOptionalSpace := optionalWhitespaceNoNewLine(openBracket)
	closedBracket := SetOfCharacters(")")
	closedBracketOptionalSpace := optionalWhitespaceNoNewLine(closedBracket)

	//name
	blockOfLetterNumberUnderscore := Range(Set(Letter, Number, underscore), 1, -1)
	letterOrUnderscore := Set(Letter, underscore)
	name := firstOrBoth(letterOrUnderscore, blockOfLetterNumberUnderscore)

	//keywords
	Func := SequenceOfCharacters("func")
	//Var := SequenceOfCharacters("var")
	//Package := SequenceOfCharacters("package")
	//Import := SequenceOfCharacters("import")

	//packageName := Label(name, "packagename")
	typeName := Label(name, "typename")
	typeNameOptionalSpace := optionalWhitespaceNoNewLine(typeName)
	variableName := Label(name, "variablename")
	returnType := Label(name, "returntype")
	//returnTypeOptionalSpace := optionalWhitespaceNoNewLine(returnType)
	functionName := Label(name, "functionName")
	functionNameOptionalSpace := optionalWhitespaceNoNewLine(functionName)

	//packageDeclaration = Sequence(Package, whitespaceNoNewLine, packageName, whitespaceAtLeastOneNewLine)

	parameter := Label(Sequence(variableName, whitespaceNoNewLine, typeName), "parameter")
	parameterOptionalSpace := optionalWhitespaceNoNewLine(parameter)
	parameterList := repeatingList(parameterOptionalSpace, commaOptionalSpace)

	functionParametersList := Sequence(openBracketOptionalSpace, parameterList, closedBracketOptionalSpace)
	functionParametersSingle := Sequence(openBracketOptionalSpace, parameterOptionalSpace, closedBracketOptionalSpace)
	functionParametersEmpty := Sequence(openBracketOptionalSpace, closedBracketOptionalSpace)

	functionParameters := Set(functionParametersList, functionParametersSingle, functionParametersEmpty)
	functionParametersOptionalSpace := optionalWhitespaceNoNewLine(functionParameters)

	returnParametersNamedList := functionParametersList
	returnParametersNamedSingle := functionParametersSingle
	returnParametersSingle := returnType
	returnParametersList := Sequence(openBracketOptionalSpace, repeatingList(typeNameOptionalSpace, commaOptionalSpace), closedBracketOptionalSpace)

	returnParameters := Set(returnParametersSingle, returnParametersList, returnParametersNamedSingle, returnParametersNamedList)
	returnParametersOptionalSpace := optionalWhitespaceNoNewLine(returnParameters)

	functionSignatureWithReturn := Sequence(Func, whitespaceNoNewLine, functionNameOptionalSpace, functionParametersOptionalSpace, returnParametersOptionalSpace)
	functionSignatureNoReturn := Sequence(Func, whitespaceNoNewLine, functionNameOptionalSpace, functionParametersOptionalSpace)

	functionSignature := Set(functionSignatureNoReturn, functionSignatureWithReturn)

	t.Run("test function signature", func(t *testing.T) {
		require.True(t, Match("func copy()", functionSignature).IsValid)
		require.True(t, Match("func copy	()", functionSignature).IsValid)
		require.True(t, Match("func copy(	)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right, float up)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) int", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (int, int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int, sedric int)", functionSignature).IsValid)
	})

	t.Run("", func(t *testing.T) {})
}
