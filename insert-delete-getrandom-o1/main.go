package main

import "math/rand"

/*
	1. Paraphrase Question
		Implement a DS that supports insert, remove, getRandom of integers in o(1) time on average.
	2. Test Cases
		Insert(1)
		GetRandom() -> 1
		Remove(1)

		Insert(1), Insert(2), Insert(3)
		GetRandom() -> ?
		Remove(1)
		GetRandom() ->
		Remove(2)
		GetRandom() -> 3

	3. Approach
		store integers in slice for random access
		use hash map to map the index they reside to
		hash int -> spot in slice for get
		when deleting always swap the last element with the deleted one, update the index

	4. Psuedo Code
		Not this time.
	5. Walk Through Test Case
		Nope

*/




type RandomizedSet struct {
	List []int
	Lookup map[int]int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{Lookup: make(map[int]int)}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rs *RandomizedSet) Insert(val int) bool {

	// check if the element already exists
	if _, ok := rs.Lookup[val]; ok {
		return false
	}

	// add the new integer to the list, make sure the lookup maps to it
	rs.List = append(rs.List, val)
	rs.Lookup[val] = len(rs.List)-1
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (rs *RandomizedSet) Remove(val int) bool {
	// check if the element doesn't already exist
	// this will also cover the list being empty
	if _, ok := rs.Lookup[val]; !ok {
		return false
	}

	// delete the last element in slice, replacing the item being deleted with the last element
	var last int
	last, rs.List = rs.List[len(rs.List)-1],rs.List[:len(rs.List)-1]

	// if val is the last element in the list we skip the swap
	if last != val {
		// replace the deleted item in the slice with the value of last
		rs.List[rs.Lookup[val]] = last
		// update the lookup so it points to the right spot
		rs.Lookup[last] = rs.Lookup[val]
	}

	// delete the look up for val
	delete(rs.Lookup,val)
	return true
}


/** Get a random element from the set. */
func (rs *RandomizedSet) GetRandom() int {
	r := rand.Intn(len(rs.List))
	return rs.List[r]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

