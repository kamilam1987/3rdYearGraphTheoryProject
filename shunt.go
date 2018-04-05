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
	pofix := []rune{} //rune-character on the screen diplayed in UTF

	//Stack which stores ooperators from the infix regular expression
	s := []rune{}

	//Loop over the input
	//range(convert into array of runes using UTF)
	for _, r := range infix {
		switch {
		//Brackets
		case r == '(':
			//open bracet at the end of the stack
			s = append(s, r)
		case r == ')':
			//Pop of the stack until finds the opening bracket
			//len(s)-1 last element in the stacl
			for s[len(s)-1] != '(' {
				pofix = append(pofix, s[len(s)-1]) //last element of s
				s = s[:len(s)-1]                   //everything in s accept the last character
			}

			s = s[:len(s)-1]
		//Current character that reads from infix is a special character
		case specials[r] > 0:
			//while stack still has an element on it and the precedence of the current character that reaads is less or = the precedence of the top element of the stcak
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				//Takes of the stack and stick into pofix
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			//Less precendence then append current character onto the stack
			s = append(s, r)
		//Not a bracket or a special character
		default:
			//Takes a character end stick and the end of output
			pofix = append(pofix, r)
		}
	}
	//If enything on the stack append to the output
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}
	return string(pofix)
} //End of intopost function

//Reference:https://rosettacode.org/wiki/Parsing/Shunting-yard_algorithm
//Function intopost which convert Maths expression to postfix regular expression,takes single argument as a string and returns string
func intopostMathsExp(infix string) string {
	//Created map with special characters into integares, keep track special characters that are allwed other then bracets
	specials := map[rune]int{'^': 4, '*': 3, '/': 3, '+': 2, '-': 2}

	//Arrray of runes, initialized an arrayof runs as empty
	pofix := []rune{} //rune-character on the screen diplayed in UTF

	//Stack which stores ooperators from the infix regular expression
	s := []rune{}

	//Loop over the input
	//range(convert into array of runes using UTF)
	for _, r := range infix {
		switch {
		//Brackets
		case r == '(':
			//open bracet at the end of the stack
			s = append(s, r)
		case r == ')':
			//Pop of the stack until finds the opening bracket
			//len(s)-1 last element in the stacl
			for s[len(s)-1] != '(' {
				pofix = append(pofix, s[len(s)-1]) //last element of s
				s = s[:len(s)-1]                   //everything in s accept the last character
			}

			s = s[:len(s)-1]
		//Current character that reads from infix is a special character
		case specials[r] > 0:
			//while stack still has an element on it and the precedence of the current character that reaads is less or = the precedence of the top element of the stcak
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				//Takes of the stack and stick into pofix
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			//Less precendence then append current character onto the stack
			s = append(s, r)
		//Not a bracket or a special character
		default:
			//Takes a character end stick and the end of output
			pofix = append(pofix, r)
		}
	}
	//If enything on the stack append to the output
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}
	return string(pofix)
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
