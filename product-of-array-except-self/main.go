package main

/*	1. Paraphrase ✓
	2. Initial Test Cases ✓
	3 Approach
		Calc the product of everything, then divd by the num at 7 at each spot to get the result.
		Track indexes with zero, if there's only 1 it's a specialcase, need to keep track of product with out it in
		If there's more than two, just need to return zero for indexes with zero nums.
	4 Pseudocode ✓
		store product in a temp var
		create boolean for more than on zero
		store location of 1 and only 1 zero
		loop through the numbers
			product = product * number
		initialize a storage slice
		loop through the numbers
			set storage slice at index to be product / nums at index
		return storage slice
	5 Walkthrough Pseudocode With Test Case
	[]int{1, 2, 3}
	product := 6
	1 = 6/1 = 6
	2 = 6/2 = 3
	3 = 6/3 = 2
	[]int{6,3,2}


*/

// productExceptSelf returns an slice with the product of every element in the slice except the ith element.
func productExceptSelf(nums []int) []int {
	// store product in a temp var
	prod := 1
	// number of zeros
	zeros := 0

	// prod without last zero index
	specprod := 1

	// loop through the numbers
	for _, val := range nums {
		// track number of zeros and index of last location, if there's only one its a special case
		if val == 0 {
			zeros++
		// skip zero value for special prod for the special case of there only being 1
		} else {
			specprod *= val
		}
		// product = product * number
		prod *= val
	}

	// initialize a storage slice
	result := make([]int, len(nums))

	// loop through the numbers
	for i, val := range nums {
		if zeros > 1  {
			result[i] = 0
		} else if zeros == 1 && val == 0 {
			// special case, exactly 1 zero
			result[i] = specprod

		} else {
			// set storage slice at index to be product / nums at index
			result[i] = prod / val
		}
	}

	return result
}