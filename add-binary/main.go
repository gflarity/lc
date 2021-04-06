package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	// setup a carry for the math
	carry := false

	result := ""
	i := 0
	for len(a)-i-1 >= 0 || len(b)-i-1 >= 0 {
		var aval byte
		// if the a string is shorter then just say it's 0
		if len(a)-1-i < 0 {
			aval = '0'
		} else {
			aval = a[len(a)-1-i]
		}

		var bval byte
		// if the b string is shorter then just say it's 0
		if len(b)-1-i < 0 {
			bval = '0'
		} else {
			bval = b[len(b)-1-i]
		}

		var reschar byte
		if carry {
			// carry + 1 + 1 -> 11
			if aval == '1' && bval == '1' {
				carry = true
				reschar = '1'
				// carry + 1 => 10
			} else if aval == '1' || bval == '1' {
				carry = true
				reschar = '0'
				// just cary
			} else {
				carry = false
				reschar = '1'
			}
			// no carry
		} else {
			if aval == '1' && bval == '1' {
				carry = true
				reschar = '0'
				// 0 + 1 => 1
			} else if aval == '1' || bval == '1' {
				carry = false
				reschar = '1'
			} else {
				carry = false
				reschar = '0'
			}
		}

		result = string(reschar) + result
		i++
	}
	if carry {
		result = "1" + result
	}
	return result
}

func main() {
	//fmt.Println(addBinary("01", "1"))
	fmt.Println(addBinary("01", "11"))
	fmt.Println(addBinary("01", "11"))
	fmt.Println(addBinary("", "11"))

	// 1000000000000000000000000000000000000000000000000000000000000000000
	fmt.Println(addBinary("100000000000000000000000000000000000000000000000000000000000000000", "100000000000000000000000000000000000000000000000000000000000000000"))

}
