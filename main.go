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
		//Shows which option was choosen
		fmt.Println("Option: ", option)
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
			fmt.Println("Error ", err.Error())
			return
		}
		//Changes from infix to postfix
		postfixExpression := shunt.Intopost(infixExpression)
		//Prints out the result
		fmt.Println("Postfix notation is: ", postfixExpression)
		//a.b.c Answer: ab.c.
		//(a.(b|d))* Answer: abd|.*
		//a.(b|d).c* Answer: abd|.c*.

	case 2:
		//Shows which option was choosen
		fmt.Println("Option: ", option)
		//Promt user to enter the postfix string
		fmt.Print("Enter postfix expression: ")
		//Reads data from user
		infixExpression, err := nfa.UserInput()
		//Error handler
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		//Promnt user to enter a regular string
		//Examples:
		//Postfix expression : ab.c*| Example of string: cccc
		//Postfix expression :ab.c*| Example of string: abc
		//Postfix expression :ab.c*| Example of string: ""

		//Promt user to enter a regular string
		fmt.Print("Enter a string to test if it matches the nfa: ")
		userString, err := nfa.UserInput()
		//Error handler
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		} //End of if
		//Checks if maches
		if nfa.Pomatch(infixExpression, userString) == true {
			//Prints result
			fmt.Println("Regular string: ", userString, " matches the expression: ", infixExpression)

			//Checks if doesn't mach
		} else if nfa.Pomatch(infixExpression, userString) == false {
			//Prints result
			fmt.Print("Regular string: ", userString, " doesn't match : ", infixExpression)
		} //End of else if

	default:
		fmt.Println("Enter one of the above options")
	} //End of switch
} //End of main
