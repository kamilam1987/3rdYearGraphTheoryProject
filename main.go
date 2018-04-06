//Author: Kamila Michel
package main

import (
	"fmt"

	"./nfa"
	"./shunt"
)

func main() {

	//Ask the user which option they like
	fmt.Println("========Menu Options=============")
	fmt.Print("Please pick option (1-3) \n 1) Infix Expression to Postfix Expresssion\n 2) Postix Regular Expresssion to NFA \n 3) Exit\n")

	//Declare variable for user option from menu
	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		//Show which option was choosen
		fmt.Println("Option:", option)
		//Promt user to enter the infix string
		fmt.Print("Enter infix expression: ")
		//Examples of Infix expressions:
		//a.b.c Answer: ab.c.
		//(a.(b|d))* Answer: abd|.*
		//a.(b|d).c* Answer: abd|.c*.

		//Reads data from user
		infixExpression, err := nfa.UserInput()
		//Error handler
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}
		//Changes from infix to postfix
		postfixExpression := shunt.Intopost(infixExpression)
		//Prints out the result
		fmt.Println("Postfix notation is:", postfixExpression)
		//a.b.c Answer: ab.c.
		//(a.(b|d))* Answer: abd|.*
		//a.(b|d).c* Answer: abd|.c*.

	default:
		fmt.Println("Enter one of the above options")
	}
} //End of main
