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
			//Pop of the stack
			frag2 := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]
			frag1 := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]

			frag1.accept.edge1 = frag2.initial

			nfastatck = append(nfastatck, &nfa{initial: frag1.initial, accept: frag2.accept})
		//When r is union character
		case '|':
			frag2 := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]
			frag1 := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastatck = append(nfastatck, &nfa{initial: &initial, accept: &accept})
		//When r is Kleene star
		case '*':
			frag := nfastatck[len(nfastatck)-1]
			nfastatck = nfastatck[:len(nfastatck)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastatck = append(nfastatck, &nfa{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastatck = append(nfastatck, &nfa{initial: &initial, accept: &accept})
		}
	}
	//Return one item at the end of stack(actual nfa that will be return)
	return nfastatck[0]
}

func main() {
	//test case
	nfa := poregtonfa("ab.c*|")
	//Prints what returns from Postfix Regular Expression to NFA
	fmt.Println(nfa)
}
