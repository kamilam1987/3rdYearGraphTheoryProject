//Author: Kamila Michel
//Source code adapted from: https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//Thompson's construction

package main

import (
	"fmt"
)

//Create state struct
type state struct {
	//Represent state that has two arrows comming from it, that are labeled by (Epsilon) or one arrow labeled by (Epsilon)
	symbol rune
	//Arrows that points to other state
	edge1 *state //One arrow
	edge2 *state
}

//(Helper struct)Keeps track of initial state and the accept state of fragment(nfa)
type nfa struct {
	initial *state
	accept  *state
}

//Function Postfix Regular Expression to NFA, takes string as input and has to return pointer to nfa
func poregtonfa(pofix string) *nfa {
	//Array of pointers to nfa's {}-one of those that are empty
	nfastatck := []*nfa{}

	//Loop true post fix reg expresion (character at the time)
	for _, r := range pofix {
		//Switch on r
		switch r {
		//When r is concatenate character
		case '.':
			//Pop of the nfa stack(last think from nfa stack)
			frag2 := nfastatck[len(nfastatck)-1]
			//To get rid of last think of the stack(up to but not including)
			nfastatck = nfastatck[:len(nfastatck)-1]
			frag1 := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]

			//Joining together (The first edge of the accept state of frag1 should point to frag2.initial)
			frag1.accept.edge1 = frag2.initial

			//Append a new pointer to nfa struct that represents this new bigger fragment)
			nfastatck = append(nfastatck, &nfa{initial: frag1.initial, accept: frag2.accept})
		//When r is union character, Pop two fragments off and push a fragment to the stack
		case '|':
			frag2 := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]
			frag1 := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]

			//Creates two new states: initial and accept
			//New state where two edges points at the initial state of the two fragments that pop off the stack
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			//New staate
			accept := state{}
			//Has two point at the accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//Push a new fragment to nfa statck
			nfastatck = append(nfastatck, &nfa{initial: &initial, accept: &accept})
		//When r is Kleene star(Pop one fragment off nfa stack)
		case '*':
			frag := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]

			//New accept state
			accept := state{}
			//New initial state which has edge 1(initial state of the pop off fragment) and edge 2(point at the new accept state)
			initial := state{edge1: frag.initial, edge2: &accept}
			//Joining together to initial state of fragment that was pop off
			frag.accept.edge1 = frag.initial
			//Joining old fragment to the accept state
			frag.accept.edge2 = &accept

			nfastatck = append(nfastatck, &nfa{initial: &initial, accept: &accept})
		default:
			//Push to the statck
			accept := state{}                           //Create a new accept state
			initial := state{symbol: r, edge1: &accept} //Create a new initial state(symbol- one of the members of the state struct)

			//Push to nfa statck
			nfastatck = append(nfastatck, &nfa{initial: &initial, accept: &accept})

		} //End of switch

	} //End of for loop

	//Return one item at the end of stack(actual nfa that will be return)
	return nfastatck[0]
} //End of poregtonfa function

func main() {
	//test case
	nfa := poregtonfa("ab.c*|")
	//Prints what returns from Postfix Regular Expression to NFA
	fmt.Println(nfa)
} //End of main
