/*  1. Paraphrase
        Given a string s, return the index of the first non-repeating character in it. Return -1 if all characters repeat at least once.

    2. Initial Test Cases
    ""
    "aa"
    "abba"
    "abc"

    3 Run through them by hand
    "aa" -> -1
    "abba" -> -1
    "abc" -> 0

    4. Psuedo code

    handle edge cases
    go through string and make a map of characters to occurrence
    go through string again and return first one with only 1 occurence
    return -1 if there weren't any

*/
func firstUniqChar(s string) int {
	// len 0 string has no characters that occur exactly once
	if len(s) == 0 {
		return -1
	}

	// len 1 string have exactly 1 character that occurs once, the only one
	if len(s) == 1 {
		return 0
	}

	// go through string and make a map of characters to occurrence
	occur := make(map[rune]int)
	for _, char := range s {
		occur[char]++
	}

	// go through string again and return first one with only 1 occurence
	for i, char := range s {
		if occur[char] == 1 {
			return i
		}
	}
	return -1
}