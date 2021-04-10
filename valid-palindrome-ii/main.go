package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	// NOTE: It would be more computationally effecient to check
	// characters/case char by char, but it would also be more complex
	// and error prone code, so I'd stick with this unless there there was
	// a requirement.
	// get rid of non alpha numerics
	var re = regexp.MustCompile("[^a-zA-Z0-9]")
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
