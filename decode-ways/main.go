package main

/*
	1. Paraphrase
		Determine the number of permutations of strings than can out of byte encoding where A -> 1... Z->26
	2. Initial Test Cases
		s := "12" -> {}string{"AB", "L"} -> 2
		s := "226" -> {}string{"BZ", "VF") -> 2
		s := "06" -> 0
		s := "0" -> 0
	3. Approach
		First pass use recursion. At each branch recursively calculate the permuations.
		Second pass unroll to avoid recursive call stack overhead once logic is solid.
		if character is 1 then either we use it as A, or we use it as 1X, recursively calculate num of combos
	4. Pseudocode
		stopping cases (string is empty -> 0 , string is only 1 character -> 1
		(string is two characters)
		if first character is 1, then we want perms of [1:] and [2:] (if we can, checking the length, as long as it's 2 there's a least 1 from this?)
		if character is 2 we want perms of [1:] + perms of [2:] (if we can, checking the length & and next chart < 7)
	5. Walk Through Psuedocode with test cases
		12.... 1 + perms of 2 = 2
		226... 1 + perms of 26 + 1 = 3
		426... 0 + perms of 26 = 0 + 2 = 2
 */
func numDecodings(s string) int {
	m := make(map[string]int)
	return numDecodingsWithMemo(m, s)
}

func numDecodingsWithMemo(m map[string]int, s string) int {
	// stopping cases (string is empty -> 0 , string is only 1 character -> 1
	if len(s) == 0 {
		return 1
	}

	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	if val, ok := m[s]; ok {
		return val
	}


	// string is two characters
	// if first character is 1, then we want perms of [1:] and [2:] (if we can, checking the length, as long as it's 2 there's a least 1 from this?)
	//if character is 2 we want perms of [1:] + perms of [2:] (if we can, checking the length & and next chart < 7)
	if s[0] == '1' || ( s[0] == '2' && s[1] != '7' && s[1] != '8' && s[1] != '9' ){
		n :=  numDecodingsWithMemo(m, s[1:]) + numDecodingsWithMemo(m, s[2:])
		m[s] = n
		return n
	}
	n := numDecodings(s[1:])
	m[s] = n
	return n
}