package main

import (
	"math"
)

/*
	1. Paraphrase
		1. Given an slice of internals [a1, z1], merge them and return a slice of non overlapping interviews that covers
		   the same ranges. Kind of like compressing, or more concisely describing the intervals.
	2. Initial Test Cases
		[[1,3],[2,6],[8,10],[15,18]] => [[1,6],[8,10],[15,18]]
		[[1,4],[4,5]] => [[1,5]]
		[[1,2],[3,4]] => [[1,2],[3,4]]
		[[1,2], [3,4], [5,6] [0,7]] => [0, 7]

		[[]] -> [[]]
	3. Approach
		Seems like we can build the "compressed' internals going interval by. Worse case you'd have no over laps
		and you'd check the n'th interval against n-1 other intervals since they don't overlap. Best case n though (one interval) It's not clear to me
		there's a way around that though since there's no overlap but you have to check for it. Maybe there's a clever
		way to do the look ups in C, or nlogn time. Could keep ranges in sorted order? Maintaining that would mean nlogn.
	4 Pseudo Code
		create compress interval result holder
		go through each of the intervals
			create temp compressed holder
			loop through all our compressed intervals
				if the compressed interval doesn't overlap, put it on our temp compressed holder
				if it does overlap skip, modify the new interval to contain both
			replace compressed interval with the tmp interval concat with the new modified interval
		return compressed intervals
	5. Walk through test cases
		[[1,2], [3,4], [5,6] [0,7]] => [0, 7]




*/
func merge(intervals [][]int) [][]int {
	// create compress interval result holder
	result := [][]int{}
	// go through each of the intervals
	for _, newival := range intervals {
		// create temp compressed holder
		tmpres := [][]int{}
		for _, oldival := range result {
			// if it does overlap skip, modify the new interval to contain both so that we pick up
			// any new mergers that result
			if merged := mergedInterval(newival, oldival); merged != nil {
				newival = merged
			} else {
				// if the compressed interval doesn't overlap, put it on our temp compressed holder
				tmpres = append(tmpres, oldival)
			}
		}

		// 	replace compressed interval with the tmp interval concat with the new modified interval
		result = append(tmpres, newival)
	}
	return result
}

// mergedInterval returns a new merged interval that spans both provided intervals, if they can't be merged
// it returns nil
func mergedInterval(a []int, b []int ) []int {
	// eq [1,3] [2, 4] or [1,2] [3, 4] or [1,2], [4,5]
	// [1,1] [2,2], [3,3]

	// if  a[1] + 1>= b[0] && a[0] + 1 <= b[0] => merge
	if a[1] >= b[0] && a[0] <= b[0] || b[1] >= a[0] && b[0] <= a[0] {

		x := math.Min(float64(a[0]), float64(b[0]))
		y := math.Max(float64(a[1]), float64(b[1]))
		return []int{int(x),int(y)}
	}

	return nil
}

