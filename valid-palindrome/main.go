package main


func validPalindrome(s string) bool {
	if !validPalindromeOneOff(s) {
		return false
	}
	return true
}

func validPalindromePerfect(s string) bool {
	i := 0
	j := len(s) - 1

	// edge case of empty string, or one letter is covered, this will be true
	for i < j {
		// atleast one is off now
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindromeOneOff(s string) bool {
	i := 0
	j := len(s) - 1

	// edge case of empty string, or one letter is covered, this will be true
	for i < j {
		// at least one is off now
		if s[i] != s[j] {
			if i+1 <= j && s[i+1] == s[j] {
				// we have a deletion
				if validPalindromePerfect(s[i+1 : j+1]) {
					return true
				}
			}
			if i <= j-1 && s[i] == s[j-1] {
				if validPalindromePerfect(s[i:j]) {
					return true
				}
			}
			// if we found a valid substring that's a perfect palinedrome we're good
			// otherwise there
			return false
		}
		i++
		j--
	}
	return true
}
