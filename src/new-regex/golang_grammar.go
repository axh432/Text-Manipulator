package new_regex

import (
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
	dot                              = SetOfCharacters(".")
	word                             = Range(Letter, 1, -1)
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



type WhitespacePattern func (expressions ...Expression) Expression

func whitespacePattern(whitespace ...Expression) WhitespacePattern {
	return func (expressions ...Expression) Expression {
		sequenceToBe := []Expression{}
		if len(whitespace) != len(expressions)-1 {
			return nil
		}
		for i := 0; i < len(whitespace); i++ {
			sequenceToBe = append(append(sequenceToBe, expressions[i]), whitespace[i])
		}
		sequenceToBe = append(sequenceToBe, expressions[len(expressions)-1])
		return Sequence(sequenceToBe...)
	}
}



//Todo: this.
//Todo: padding = delimiter and whitespace
func RepeatingList(listItem Expression, padding Expression) Expression {
	return Sequence(Range(Sequence(listItem, padding), 1, -1), listItem)
}

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
