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
	boolValue                        = Set(SequenceOfCharacters("true"), SequenceOfCharacters("false"))
	integerValue                     = Range(Number, 1, -1)
	listOfIntegerValues              = Sequence(Range(Sequence(integerValue, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), integerValue)

	//name
	letterNumberUnderscoreBlock         = Range(Set(Letter, Number, underscore), 1, -1)
	optionalLetterNumberUnderscoreBlock = Range(Set(Letter, Number, underscore), 0, -1)
	letterOrUnderscore                  = Set(Letter, underscore)
	name                                = Sequence(letterOrUnderscore, optionalLetterNumberUnderscoreBlock)

	typeName     = Label(name, "typename")
	variableName = Label(name, "variablename")
	returnType   = Label(name, "returntype")
	functionName = Label(name, "functionName")

	//Import
	Import                  = SequenceOfCharacters("import")
	importNameWithSpecifier = Sequence(SetOfCharacters("_."), optionalWhitespaceNoNewLineBlock, String)
	importNameNoSpecifier   = String
	importName              = Set(importNameWithSpecifier, importNameNoSpecifier)
	importMultiple          = Sequence(Range(Sequence(importName, whitespaceAtLeastOneNewLineBlock), 1, -1), importName)
	importBoundedMultiple   = Sequence(openBracket, optionalWhitespaceBlock, importMultiple, optionalWhitespaceBlock, closedBracket)
	importBoundedSingle     = Sequence(openBracket, optionalWhitespaceBlock, importName, optionalWhitespaceBlock, closedBracket)
	importBoundedEmpty      = Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	importBoundedAll        = Set(importBoundedMultiple, importBoundedSingle, importBoundedEmpty)
	importSingle            = importName
	importDeclaration       = Sequence(Import, optionalWhitespaceBlock, Set(importBoundedAll, importSingle))

	//Function Signature
	Func                        = SequenceOfCharacters("func")
	parameter                   = Label(Sequence(variableName, whitespaceNoNewLineBlock, typeName), "parameter")
	functionParameters          = FunctionParameterCodeBlock(openBracket, parameter, comma, closedBracket)
	returnParametersNamed       = functionParameters
	returnParametersSingle      = returnType
	returnParametersList        = Sequence(Range(Sequence(returnType, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), optionalWhitespaceBlock, returnType)
	returnParametersBoundedList = Sequence(openBracket, optionalWhitespaceBlock, returnParametersList, optionalWhitespaceNoNewLineBlock, closedBracket)
	returnParameters            = Set(returnParametersSingle, returnParametersBoundedList, returnParametersNamed)
	optionalReturnParameters    = Range(returnParameters, 0, 1)
	functionSignature           = Sequence(Func, whitespaceNoNewLineBlock, functionName, optionalWhitespaceNoNewLineBlock, functionParameters, optionalWhitespaceNoNewLineBlock, optionalReturnParameters)

	//function call
	functionCallParametersMultiple        = Sequence(Range(Sequence(variableName, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), variableName)
	functionCallParametersBoundedMultiple = Sequence(openBracket, optionalWhitespaceBlock, functionCallParametersMultiple, optionalWhitespaceNoNewLineBlock, closedBracket)
	functionCallParametersBoundedSingle   = Sequence(openBracket, optionalWhitespaceBlock, variableName, optionalWhitespaceNoNewLineBlock, closedBracket)
	functionCallParametersBoundedEmpty    = Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	functionCallParametersBoundedAll      = Set(functionCallParametersBoundedMultiple, functionCallParametersBoundedSingle, functionCallParametersBoundedEmpty)
	optionalPackageName                   = Range(Sequence(packageName, optionalWhitespaceNoNewLineBlock, dot, optionalWhitespaceBlock), 0, 1)
	functionCall                          = Sequence(optionalPackageName, functionName, optionalWhitespaceNoNewLineBlock, functionCallParametersBoundedAll)

	//Var Assign Statement
	Var                   = SequenceOfCharacters("var")
	varAssignmentOperator = SetOfCharacters("=")
	valuePossibilities    = Set(String, boolValue, integerValue, listOfIntegerValues)
	optionalTypeName      = Range(typeName, 0, 1)
	varNameList           = Sequence(Range(Sequence(variableName, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), variableName)
	varNames              = Set(varNameList, variableName)
	varAssignStatement    = Sequence(Var, optionalWhitespaceBlock, varNames, optionalWhitespaceNoNewLineBlock, optionalTypeName, optionalWhitespaceNoNewLineBlock, varAssignmentOperator, optionalWhitespaceBlock, valuePossibilities)

	//Assign Statement
	assignmentOperator = SequenceOfCharacters(":=")
	assignStatement    = Sequence(varNames, optionalWhitespaceNoNewLineBlock, assignmentOperator, optionalWhitespaceBlock, valuePossibilities)

	//Package
	Package            = SequenceOfCharacters("package")
	packageName        = Label(name, "packagename")
	packageDeclaration = Sequence(Package, whitespaceNoNewLineBlock, packageName)
)

type WhitespacePattern func(expressions ...Expression) Expression

func whitespacePattern(whitespace ...Expression) WhitespacePattern {
	return func(expressions ...Expression) Expression {
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

func MultiLineRepeatingList(listItem Expression, delimiter Expression) Expression {
	return Sequence(Range(Sequence(listItem, optionalWhitespaceNoNewLineBlock, delimiter, optionalWhitespaceBlock), 1, -1), optionalWhitespaceBlock, listItem)
}

func MultiLineRepeatingListNoDelimiter(listItem Expression) Expression {
	return Sequence(Range(Sequence(listItem, whitespaceAtLeastOneNewLineBlock), 1, -1), listItem)
}

func BoundedMultiLineRepeatingListNoDelimiter(start Expression, listItem Expression, end Expression) Expression {
	return Sequence(start, optionalWhitespaceBlock, MultiLineRepeatingListNoDelimiter(listItem), optionalWhitespaceBlock, end)
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
