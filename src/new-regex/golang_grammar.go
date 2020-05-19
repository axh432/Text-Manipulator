package new_regex

import (
	"errors"
	_ "fmt"
	_ "strings"
)

var (
	//Primitives
	whitespaceBlock                  = Range(Whitespace, 1, -1)
	optionalWhitespaceBlock          = Range(Whitespace, 0, -1)
	whitespaceNoNewLine              = SetOfCharacters(" \t")
	whitespaceNoNewLineBlock         = Range(whitespaceNoNewLine, 1, -1)
	optionalWhitespaceNoNewLineBlock = Range(SetOfCharacters(" \t"), 0, -1)
	newline                          = SequenceOfCharacters("\n")
	whitespaceAtLeastOneNewLineBlock = Sequence(optionalWhitespaceNoNewLineBlock, newline, optionalWhitespaceBlock)
	underscore                       = SetOfCharacters("_")
	comma                            = SetOfCharacters(",")
	openBracket                      = SetOfCharacters("(")
	closedBracket                    = SetOfCharacters(")")
	quote                            = SetOfCharacters(`"`)
	String                           = Sequence(quote, Range(Set(SetOfNotCharacters(`"`), SequenceOfCharacters(`\"`)), 1, -1), quote)

	//name
	letterNumberUnderscoreBlock         = Range(Set(Letter, Number, underscore), 1, -1)
	optionalLetterNumberUnderscoreBlock = Range(Set(Letter, Number, underscore), 0, -1)
	letterOrUnderscore                  = Set(Letter, underscore)
	name                                = Sequence(letterOrUnderscore, optionalLetterNumberUnderscoreBlock)

	//keywords
	Func    = SequenceOfCharacters("func")
	Var     = SequenceOfCharacters("var")
	Package = SequenceOfCharacters("package")
	Import  = SequenceOfCharacters("import")

	packageName  = Label(name, "packagename")
	typeName     = Label(name, "typename")
	variableName = Label(name, "variablename")
	returnType   = Label(name, "returntype")
	functionName = Label(name, "functionName")

	importNameWithSpecifier = Sequence(SetOfCharacters("_."), optionalWhitespaceNoNewLineBlock, String)
	importNameNoSpecifier   = String

	importName        = Set(importNameWithSpecifier, importNameNoSpecifier)
	importBlock       = MultiLineCodeBlockNoDelimiter(openBracket, importName, closedBracket)
	importSingle      = String
	importDeclaration = Sequence(Import, optionalWhitespaceBlock, Set(importBlock, importSingle))

	packageDeclaration = Sequence(Package, whitespaceNoNewLineBlock, packageName)

	parameter = Label(Sequence(variableName, whitespaceNoNewLineBlock, typeName), "parameter")

	functionParameters = FunctionParameterCodeBlock(openBracket, parameter, comma, closedBracket)

	returnParametersNamed  = functionParameters
	returnParametersSingle = returnType
	returnParametersList   = FunctionParameterList(openBracket, returnType, comma, closedBracket)

	returnParameters         = Set(returnParametersSingle, returnParametersList, returnParametersNamed)
	optionalReturnParameters = Range(returnParameters, 0, 1)

	functionSignature = Sequence(Func, whitespaceNoNewLineBlock, functionName, optionalWhitespaceNoNewLineBlock, functionParameters, optionalWhitespaceNoNewLineBlock, optionalReturnParameters)
)

func interlaceWhitespace(tokens []Expression, whitespace []Expression) (Expression, error) {

	tokenLen := len(tokens)
	whitespaceLen := len(whitespace)
	sequenceToBe := []Expression{}

	if tokenLen == 0 || whitespaceLen == 0 {
		return nil, errors.New("one or both of the slices provided has a length of 0")
	} else if whitespaceLen == tokenLen {
		for i, _ := range tokens {
			sequenceToBe = append(append(sequenceToBe, tokens[i]), whitespace[i])
		}
	} else if whitespaceLen-1 == tokenLen {
		for i := 0; i < tokenLen-1; i++ {
			sequenceToBe = append(append(sequenceToBe, tokens[i]), whitespace[i])
		}
		sequenceToBe = append(sequenceToBe, tokens[tokenLen-1])
	} else {
		return nil, errors.New("whitespace slice has to be either equal or one less than the token slice")
	}

	return Sequence(sequenceToBe...), nil
}

func listDelimiterStartAndEnd(start Expression, listItem Expression, delimiter Expression, end Expression, whitespace []Expression) Expression{
	itemAndDelimiter := Sequence(listItem, whitespace[1], delimiter)
	list := Sequence(Range(Sequence(itemAndDelimiter, whitespace[2]), 1, -1), listItem)
	multiple := Sequence(start, whitespace[0], list, whitespace[3], end)
	single := Sequence(start, whitespace[0], listItem, whitespace[3], end)
	empty := Sequence(start, whitespace[3], end)
	return Set(single, multiple, empty)
}

//Todo: with whitespace, with delimiter, with start and end expression, with single and with empty
func interlaceWhitespaceList(listItem Expression, delimiter Expression, whitespace []Expression) Expression {
	itemAndDelimiter := createItemAndDelimiter(listItem, delimiter, whitespace)
	return createList(itemAndDelimiter, listItem, whitespace)
}

func createItemAndDelimiter(listItem Expression, delimiter Expression, whitespace []Expression) Expression {
	if whitespace == nil || len(whitespace) == 0 {
		if delimiter == nil {
			return listItem
		}
		return Sequence(listItem, delimiter)
	}

	if delimiter == nil {
		return Sequence(listItem, whitespace[0])
	}

	return Sequence(listItem, whitespace[0], delimiter)
}

func createList(listItemAndDelimiter Expression, listItem Expression, whitespace []Expression) Expression {
	if whitespace == nil || len(whitespace) == 0 {
		return Sequence(Range(listItemAndDelimiter, 1, -1), listItem)
	}else if whitespace == nil || len(whitespace) == 1 {
		return Sequence(Range(Sequence(listItemAndDelimiter, whitespace[0]), 1, -1), listItem)
	}
	return Sequence(Range(Sequence(listItemAndDelimiter, whitespace[1]), 1, -1), listItem)
}


/*func RepeatingList(listItem Expression, delimiter Expression) Expression {
	return Sequence(Range(Sequence(listItem, delimiter), 1, -1), listItem)
}*/

func SingleLineRepeatingList(listItem Expression, delimiter Expression) Expression {
	return Sequence(Range(Sequence(listItem, optionalWhitespaceNoNewLineBlock, delimiter, optionalWhitespaceNoNewLineBlock), 1, -1), optionalWhitespaceNoNewLineBlock, listItem)
}

func SingleLineRepeatingListNoDelimiter(listItem Expression) Expression {
	return Sequence(Range(Sequence(listItem, whitespaceNoNewLineBlock), 1, -1), listItem)
}

func MultiLineRepeatingList(listItem Expression, delimiter Expression) Expression {
	return Sequence(Range(Sequence(listItem, optionalWhitespaceNoNewLineBlock, delimiter, optionalWhitespaceBlock), 1, -1), optionalWhitespaceBlock, listItem)
}

func MultiLineRepeatingListNoDelimiter(listItem Expression) Expression {
	return Sequence(Range(Sequence(listItem, whitespaceAtLeastOneNewLineBlock), 1, -1), listItem)
}

func BoundedSingleLineRepeatingList(start Expression, listItem Expression, delimiter Expression, end Expression) Expression {
	return Sequence(start, optionalWhitespaceNoNewLineBlock, SingleLineRepeatingList(listItem, delimiter), optionalWhitespaceNoNewLineBlock, end)
}

func BoundedMultiLineRepeatingList(start Expression, listItem Expression, delimiter Expression, end Expression) Expression {
	return Sequence(start, optionalWhitespaceBlock, MultiLineRepeatingList(listItem, delimiter), optionalWhitespaceBlock, end)
}

func BoundedMultiLineRepeatingListNoDelimiter(start Expression, listItem Expression, end Expression) Expression {
	return Sequence(start, optionalWhitespaceBlock, MultiLineRepeatingListNoDelimiter(listItem), optionalWhitespaceBlock, end)
}

func SingleLineCodeBlock(start Expression, listItem Expression, delimiter Expression, end Expression) Expression {
	list := BoundedSingleLineRepeatingList(start, listItem, delimiter, end)
	single := Sequence(start, optionalWhitespaceNoNewLineBlock, listItem, optionalWhitespaceNoNewLineBlock, end)
	empty := Sequence(start, optionalWhitespaceNoNewLineBlock, end)
	return Set(list, single, empty)
}

func MultiLineCodeBlock(start Expression, listItem Expression, delimiter Expression, end Expression) Expression {
	list := BoundedMultiLineRepeatingList(start, listItem, delimiter, end)
	single := Sequence(start, optionalWhitespaceBlock, listItem, optionalWhitespaceBlock, end)
	empty := Sequence(start, optionalWhitespaceBlock, end)
	return Set(list, single, empty)
}

func MultiLineCodeBlockNoDelimiter(start Expression, listItem Expression, end Expression) Expression {
	list := BoundedMultiLineRepeatingListNoDelimiter(start, listItem, end)
	single := Sequence(start, optionalWhitespaceBlock, listItem, optionalWhitespaceBlock, end)
	empty := Sequence(start, optionalWhitespaceBlock, end)
	return Set(list, single, empty)
}

func FunctionParameterList(start Expression, listItem Expression, delimiter Expression, end Expression) Expression {
	return Sequence(start, optionalWhitespaceBlock, MultiLineRepeatingList(listItem, delimiter), optionalWhitespaceNoNewLineBlock, end)
}

func FunctionParameterCodeBlock(start Expression, listItem Expression, delimiter Expression, end Expression) Expression {
	list := FunctionParameterList(start, listItem, delimiter, end)
	single := Sequence(start, optionalWhitespaceBlock, listItem, optionalWhitespaceBlock, end)
	empty := Sequence(start, optionalWhitespaceBlock, end)
	return Set(list, single, empty)
}
