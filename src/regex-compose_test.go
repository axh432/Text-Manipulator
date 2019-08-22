package main

import (
	"fmt"
	. "regexp"
	"strings"
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

type Keyword struct {}
type Name struct {}
type Symbol struct {}

type Package struct {
	Name string
	subPackage *Package
}

type ImportStatement struct {
	ImportKeyWord string
	Package *Package
}

func parseImportStatement(begin int, splitStrings []string) (int, ImportStatement){

	for i := begin; i < len(splitStrings); i++ {

		token := splitStrings[i]

 		if strings.Contains(token, "."){

		}else if strings.Contains(token, ";"){

		}
	}

}

func Test_parse(t *testing.T){

	re := MustCompile(`[[:punct:]]|[[:space:]]`)

	splitStrings := splitKeepDelimiter(re, codeReference)

	for i := 0; i < len(splitStrings); i++ {

		token := splitStrings[i]

		if strings.Contains(token, "import"){
		}

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