package main

import "fmt"

func main() {

	fmt.Println(isValid("[{({}[]())}]"))
}

func isValid(s string) bool {
	// empty string is still a valid and balanced technically
	if len(s) == 0 {
		return true
	}
	// since the string only contains parens, an odd length string must be
	// unbalanced
	if len(s)%2 == 1 {
		return false
	}

	var stack []rune
	for _, next := range s {
		// if the stack is empty, we should be pushing on to it, nothing to balance
		if len(stack) == 0 {

			// stack is empty and we've been given a closing paren, s is invalid
			if isClosing(next) {
				return false
			}

			stack = push(stack, next)
			continue
		}

		// we're going to want to peak
		peaked := peak(stack)
		//		s := string(stack)

		// if the next rune is a close rune, then the peaked rune better be the opposite
		if isClosing(next) && peaked == opposite(next) {
			// now we can pop off the stack and clean up
			stack, _ = pop(stack)
			continue
		} else if isClosing(next) {
			// if it's any other closing rune, this is an unbalanced string
			return false
		}

		// it's an opening rune then pop on to the stack
		stack = push(stack, next)
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func printStack(stack []rune) {
	fmt.Println(string(stack))
}

func opposite(r rune) rune {
	if r == '(' {
		return ')'
	}
	if r == '[' {
		return ']'
	}
	if r == '{' {
		return '}'
	}
	if r == '}' {
		return '{'
	}
	if r == ')' {
		return '('
	}
	// only option left
	return '['
}

// isOpening checks if a rune is an opening paren of some sort
func isOpening(r rune) bool {
	if r == '(' || r == '[' || r == '[' {
		return true
	}
	return false
}

// isClosing check if a rune is a closing paren of some sort
func isClosing(r rune) bool {
	if r == ')' || r == ']' || r == '}' {
		return true
	}
	return false
}

// push convenience function for making a slice look more like a stack
func push(stack []rune, r rune) []rune {
	return append(stack, r)
}

// pop simulates poping an element form the slice, it returns a new slice
// with stack and the rune that was popped.
func pop(stack []rune) ([]rune, rune) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

// peak is just a convience abstraction to make the code easier to read
func peak(stack []rune) rune {
	return stack[len(stack)-1]
}
