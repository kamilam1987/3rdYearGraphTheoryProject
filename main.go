//Author: Kamila Michel
//source code adapted from: https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e

package main

//imports
import (
	"fmt"
)

//Function intopost which convert regular expression to postfix regular expression,takes single argument as a string and returns string
func intopost(infix string) string {
	//Created map with special characters into integares, keep track special characters that are allwed other then bracets
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	//Arrray of runes, initialized an arrayof runs as empty
	postfix := []rune{} //rune-character on the screen diplayed in UTF

	//Stack which stores ooperators from the infix regular expression
	s := []rune{}

	//Loop over the input
	for _, r := range infix {

	}

	return string(postfix)
} //End of intopost function

//Main function contains regular expression written in infix notation
func main() {
	//Answer: ab.c*.
	fmt.Println("Infix:  ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix:  ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	//Answer: abd|.c*.
	fmt.Println("Infix:  ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix:  ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))
} //End of main function
