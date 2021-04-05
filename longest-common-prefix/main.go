package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	// special cases up front, empty or 1 string slice
	if len(strs) == 0 {
		// nothing to check prefixes on
		return ""
	} else if len(strs) == 1 {
		// only one string, it IS the prefix
		return strs[0]
	}

	prefix := ""

	// loop on through characters until until there's a
	// difference or a string ends
	i := 0
	for {
		// loop through each string, this gives us a rune
		first := strs[0]

		// gone past the bounds of first word, we're done
		//
		if len(first) < i+1 {
			return prefix
		}

		// the char we're looking for
		char := first[i]

		// if at any point we don't find it, we're done with what
		// we've got already
		for j := 1; j < len(strs); j++ {
			str := strs[j]
			if len(str) < i+1 {
				return prefix
			}

			if str[i] != char {
				return prefix
			}
		}

		// increase the prefix
		prefix += string(char)

		// next char
		i++

	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"one", "one hundred", "once"}))
	fmt.Println(longestCommonPrefix([]string{"one", "two", "three"}))

}
