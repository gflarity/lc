package main

import "container/heap"

/*
	1. Paraphrase
		Return the kth most frequent integer in the slice.
	2. Initial Testcases
		nums = []int{1,1,1,2,2,3}, k = 2
		nums = []int{1}, k = 1
	3. Approach
		Wrap integers in a struct that contains frequency.
		Keep a map of integer to their struct
		Use a heap as a max equeue, pop from it k times at the end
		could fix heap on the fly as we walk, probably about the same as just doing in the beginning and much
		easier to follow/read.
	4. Pseudo Code
	5  Walkthrough psuedo code in case

 */


func topKFrequent(nums []int, k int) []int {
	q := NewMaxQueue(nums)
	var topk []int
	for i:=0; i < k; i++ {
		topk = append(topk, heap.Pop(q).(*FreqTracker).val)
	}
	return topk
}

type FreqTracker struct {
	val int
	freq int
}

type MaxQueue struct {
	d []*FreqTracker
}

func NewMaxQueue(nums []int) *MaxQueue {
	m := make(map[int]*FreqTracker)
	var fnums []*FreqTracker
	for _, val := range nums {
		if _, ok := m[val]; !ok {
			// new val being tracking, initialize
			ft := &FreqTracker{freq: 1, val: val}
			m[val] = ft
			fnums = append(fnums, ft)
			continue
		} else {
			// bump the frequency
			f := m[val]
			f.freq++
		}
	}

	// create and initialize the max queue
	q := &MaxQueue{d: fnums}
	heap.Init(q)
	return q
}

func (q *MaxQueue) Len() int {
	return len(q.d)
}

func (q *MaxQueue) Less(i, j int) bool {
	if q.d[i].freq > q.d[j].freq {
		return true
	}
	return false
}

func (q *MaxQueue) Swap(i,j int) {
	tmp := q.d[i]
	q.d[i] = q.d[j]
	q.d[j] = tmp
}

func (q *MaxQueue) Push(x interface{}) {
	q.d = append(q.d, x.(*FreqTracker))
}

func (q *MaxQueue) Pop() interface{}   {
	var last interface{}
	last, q.d = q.d[len(q.d)-1], q.d[:len(q.d)-1]
	return last
}
