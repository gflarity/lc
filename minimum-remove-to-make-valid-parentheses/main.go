package main
/*
	1. Paraphrase
		Remove the minimal amount of parenthesis so that the string is valid (balanced parens).
	2. Test Cases
		Done
	3. Approach
		Use stack to keep track of balance. Remove every bracket that is out of balance.
	4 Pseudo Code
		init stack
		loop through every character
			if the character is ( push onto stack and add to new string
			if the character is )  and stack is empty we need to remove this
			if the character is ) and the stack is not empty
				pop from stack and add to new string
			ignore any other character

		if the stack still has elements, loop through the elments backwards
			if character is ) push onto stack
			if character is ( and the stack is not empty, pop from it
			if chracter is ( and the stack is empty we need to remove this one
			ignore any other character
	5. Walk through Psuedo Code with a good test case.
		((())
		s: (
		s: ((
		s: (((
		s: ((
		s: ( too many

		s: )
		s: ))
		s: )
		s:
		remove
		done

 */


func minRemoveToMakeValid(s string) string {
	ns := ""
	stack := []rune{}
	// loop through every character
	for _, r := range s {
		// if the character is ( push onto stack and add to new string
		if r == rune('(') {
			stack = append(stack, r)
			ns += string(r)
		} else if r == rune(')') && len(stack) != 0 {
			// if the character is ) and the stack is not empty
			// pop from stack and add to new string
			stack = stack[:len(stack)-1]
			ns += string(r)
		} else if r == rune(')') && len(stack) == 0 {
			// if the character is )  and stack is empty we need to remove this
			continue
		} else {
			// else this isn't a bracket, include but don't touch stack
			ns += string(r)
		}
	}
	if len(stack) == 0 {
		return ns
	}
	s = ns
	ns = ""
	stack = []rune{}
	// if the stack still has elements, look through the elments backwards
	for i := len(s) - 1 ; i > -1; i-- {
		r := rune(s[i])
		// if character is ) push onto stack
		if r == rune(')') {
			stack = append(stack, r)
			ns = string(r) + ns
		} else if r == rune('(') && len(stack) != 0 {
			// if character is ( and the stack is not empty, pop from it
			ns = string(r) + ns
			stack = stack[:len(stack)-1]
		} else if r == rune('(') && len(stack) == 0 {
			continue
			// if chracter is ( and the stack is empty we need to remove this one
		} else {
			// ignore any other character but include it
			ns = string(r) + ns
		}
	}
	return ns
}
