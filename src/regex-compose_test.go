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
	Text string
	Tags []string
}


func tag(regexPattern string, tag string, token *Token){
	regexExp := MustCompile(regexPattern)

	if regexExp.MatchString(token.Text) {
		token.Tags = append(token.Tags, tag)
	}
}

func tagForWhitespace(token *Token){
	tag(`\s+`, "whitespace", token)
}

func tagForKeyword(token *Token){
	tag(`import|class|public|static|void|main|if|int|String|new|else`, "keyword", token)
}

func tagForSymbol(token *Token){
	tag(`[[:punct:]]`, "symbol", token)
}

func tagForWord(token *Token){
	tag(`\w+`, "word", token)
}

func tagForNumber(token *Token){
	tag(`\d+`, "number", token)
}

func tagForEmptyString(token *Token){
	tag(`^$`, "empty string", token)
}

func tagToken(token *Token){

	tagForWhitespace(token)
	tagForKeyword(token)
	tagForSymbol(token)
	tagForWord(token)
	tagForNumber(token)
	tagForEmptyString(token)

}

func Test_parse(t *testing.T){

	tokens := []Token{}

	re := MustCompile(`[[:punct:]]|[[:space:]]`)

	splitStrings := splitKeepDelimiter(re, codeReference)

	for _, rawToken := range splitStrings {

		token := Token{ rawToken, []string{}  }

		tagToken(&token)

		tokens = append(tokens, token)

		fmt.Printf("text: %q, tags: %v\n", token.Text, token.Tags)
	}

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

	exampleBuild := fileAsSection.find(MustCompile(QuoteMeta(`        stage('Example Build') {
            agent { docker 'maven:3-alpine' } 
            steps {
                echo 'Hello, Maven'
                sh 'mvn --version'
            }
        }`)))


	if exampleBuild.isEmpty() {
		t.Errorf("Example build could not be found. Did you forget to QuoteMeta?")
		return
	}

	fmt.Println(exampleBuild.toString())

}