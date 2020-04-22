package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_match(t *testing.T) {

	t.Run("match the string", func(t *testing.T) {
		stringToMatch := "voidvoid (l r i){+}"

		exp := Sequence(Range(SequenceOfCharacters("void"), 2, 2), Whitespace, SequenceOfCharacters("(l r i){+}"))

		result := Match2(stringToMatch, exp)

		require.True(t, result.isValid)
		println(result.toString())
	})

/*	t.Run("match the string", func(t *testing.T) {
		stringToMatch := "public void myFunction(int left, int right) int"

		type Parameter struct {
			Type string
			Name string
		}

		type FunctionSignature struct {
			Keywords []string
			Name []string
			Parameters []Parameter
			ReturnType string
		}

		//Primitives
		underscore := SetOfCharacters("_")
		comma := SetOfCharacters(",")
		openbracket := SetOfCharacters("(")
		closedBracket := SetOfCharacters(")")

		whitespaceNoNewLine := SetOfCharacters(" \t")
		optionalWhitespaceNoNewLine := Range(whitespaceNoNewLine, 0, -1)

		name := Sequence(Letter, Set(Letter, Number, underscore))

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

		require.True(t,
			Match(stringToMatch, functionSignature))
	})*/

}
