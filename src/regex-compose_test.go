package main

import (
	"fmt"
	. "regexp"
	"testing"
)

var codeReference string = `
import java.util.Scanner;

class CheckEvenOdd
{
  public static void main(String args[])
  {
    int num;
    System.out.println("Enter an Integer number:");

    //The input provided by user is stored in num
    Scanner input = new Scanner(System.in);
    num = input.nextInt();

    /* If number is divisible by 2 then it's an even number
     * else odd number*/
    if ( num % 2 == 0 )
        System.out.println("Entered number is even");
     else
        System.out.println("Entered number is odd");
  }
}
`

type Token struct {
	text string
	tags []string
}


//TODO: It would be better to group things into basic groups: keywords, names, symbols. Then look for more identifying features.

func tagToken(rawToken string) Token {

	token := Token{ rawToken, []string{}  }

	tagIfWhitespace(rawToken)

}

func Test_parse(t *testing.T){

	//tokens := []Token{}

	re := MustCompile(`[[:punct:]]|[[:space:]]`)

	splitStrings := splitKeepDelimiter(re, codeReference)

	for _, token := range splitStrings {
		fmt.Println(token)
	}


	//every line needs to start with importkeyword
	//every line needs to have a package name followed by the import keyword
	//every line needs to have a line terminator
	//every line can optionally have a further accessorDot followed by another package name

	//legalSyntax := "ImportStatement:ImportKeyword,NameRegex,{AccessorDot,NameRegex}*?,LineTerminator"

}


func Test_splitKeepDelimiter(t *testing.T){

	re := MustCompile(`[[:punct:]]|[[:space:]]`)

	splitStrings := splitKeepDelimiter(re, codeReference)

	for _, token := range splitStrings {
		fmt.Println(token)
	}


}

func Test_findStringIndex(t *testing.T){

	fileAsSection := createSectionFromString(codeSection)

	exampleBuild, err := fileAsSection.find(MustCompile(QuoteMeta(`        stage('Example Build') {
            agent { docker 'maven:3-alpine' } 
            steps {
                echo 'Hello, Maven'
                sh 'mvn --version'
            }
        }`)))


	if err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Println(exampleBuild.toString())

}