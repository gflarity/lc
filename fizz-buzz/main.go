package main

import "fmt"

// approach
// x 1 - use % to determine if each number is a multiple as we go
// 		for really large ns, this could be expensive
// x 2 - pre compute all the multiples of 3 and 5
// o(n) compute
// o(n) memory
// * 3 - as we go keep two counters, one for 3s one for 5s, reset them 0
// whenever they increment to 3 and 5, reset to 0
//   - o(c) memory
//   - o(n) compute
func fizzBuzz(n int) []string {
	// note that zero is a multiple of 3 and 5 (0*3*5)
	three := 0
	five := 0

	// inclusive of n
	var results []string
	for i := 1; i < n+1; i++ {
		var rep string
		three++
		if three == 3 {
			rep += fmt.Sprint("Fizz")
			three = 0
		}
		five++
		if five == 5 {
			five = 0
			rep += fmt.Sprint("Buzz")
		}

		if three != 0 && five != 0 {
			rep += fmt.Sprintf("%d", i)
		}
		results = append(results, rep)
	}
	return results
}

func main() {
	// initial/obvious test cases
	fmt.Println(fizzBuzz(0))
	fmt.Println(fizzBuzz(1))
	fmt.Println(fizzBuzz(-1))
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))

	// 3. test case second pass
	// 4. implementation
}
