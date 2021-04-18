package main

/*
	1. Paraphrase
		Given a vector as a slice, store it sparsely and support a cross product.
	2. Initial test cases
 		[0,1,0,0,0,0,0,3] * [0,2,0,0,0,0,0,0] = 2
		[0,0,0,0,0,0,0,0] * [1,2,3,4,5,7,8] = 0
		[1,0,0,0,0,0,0,1] * [1,2,3,4,5,7,8] = 9
	3. Approach
		store the vectors as pairs [index, value] and then it's zero for everything else
		multiple them to together by incrementing two counters, it's only non zero if the indexes match
	4. Pseudo Code
		constructor
		create two dim vector collection
		loop through each index,
			if vector at index is non-zero store the index,value pair in the collection
		return collection along with the length of the vector

		dotProduct
		i, j, sum = 0, 0, 0
		loop until i > lenth of v1 ||  j > length of v2
			if v1[i][0] == v2[j][0]
				sum = sum + v1[i][1]*v2[j][1]
				i++
				j++
			else if v1[i] < v2[j]
				i++
			else if v2[j] < v1[i]
				j++
		return sum

	5. Run test case through pseudo code
		[0,1,0,0,0,0,0,3] * [0,2,0,0,0,0,0,0] = 2
		[[1,1][7,3]] [[1,2]]

		i, j, sum = 0,0,0
			1 = 1
				sum = sum + 1*2 = 2

*/
type SparseVector struct {
	rep [][2]int
}

func Constructor(nums []int) SparseVector {
	// create two dim vector collection
	var rep [][2]int

	// loop through each index
	for i, val := range nums {
		// if vector at index is non-zero store the index,value pair in the collection
		if val != 0 {
			rep = append(rep, [2]int{i, val})
		}
	}

	//return collection along with the length of the vector
	return SparseVector{
		rep:  rep,
	}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	// i, j, sum = 0, 0, 0
	var i, j, sum int

	// loop until i > lenth of v1 ||  j > length of v2
	for i < len(this.rep) && j < len(vec.rep) {
		// if v1[i][0] == v2[j][0]
		if this.rep[i][0] == vec.rep[j][0] {
			sum += this.rep[i][1]*vec.rep[j][1]
			i++
			j++
		} else if this.rep[i][0] < vec.rep[j][0] {
			i++
		} else {
			//	this.rep[i][0] > vec.rep[j][0]
			j++
		}
	}
	return sum
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */