package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {

	// walk right to left and do math like we used to in school
	i1 := len(num1) - 1
	i2 := len(num2) - 1
	carry := 0
	result := ""
	for i1 > -1 || i2 > -1 {
		nc1 := 0
		if i1 >= 0 {
			nc1 = parseInt(num1[i1])
			i1--
		}
		nc2 := 0
		if i2 >= 0 {
			nc2 = parseInt(num2[i2])
			i2--
		}

		tc := nc1 + nc2
		if carry != 0 {
			tc = tc + carry
		}

		// biggest number is 9+9 = 18
		carry = 0
		if tc > 10 {
			carry = 1
			tc = tc % 10
		}

		// it would be more effecient to just create a byte array then reverse
		// then just reverse it in the end, but this is good enough for first pass
		result = fmt.Sprintf("%d%s", tc, result)
	}
	return result
}

func parseInt(b byte) int {
	// this can be optimized with a big switch/if block
	i, _ := strconv.ParseInt(string(b), 10, 64)
	return int(i)
}

func main() {
	fmt.Println(addStrings("11", "123"))
}
