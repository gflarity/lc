package main

import "container/heap"

/*
	1. Paraphase Question
		Given a slice and value k, find the kth biggest element in the slice without sorting it (ie it should perform
		better).
	2. Intiial Test Cases
		nums = []int{3,2,1,5,6,4},  k = 2 -> 5
		nums = []int{3,2,3,1,2,4,5,5,6}, k = 4 -> 4
		nums = []int{1}, k = 1 -> 1
	3 Approach
		Build a max heap of size k ish
		Populate the heap
		then pull from the heap k times to get the number
	4 Pseudo Code
		Insert
			append item onto end of queue
			while the parent is < then  item
				swap item with parent
				update parent
		Delete
			swap last element with item beeing deleted then delete
			while left child or right right of current last item home is bigger than item
				swap with largest child
				update current home

	5 Walk through test case with psuedo code
 */

func findKthLargest(nums []int, k int) int {
	q := NewMaxQueue(nums)
	heap.Init(q)
	var kth int
	for i := 0; i < k; i++ {
		kth = heap.Pop(q).(int)
	}
	return kth
}

type MaxQueue struct {
	d []int
}

func NewMaxQueue(nums []int) *MaxQueue {
	return &MaxQueue{
		d: nums,
	}
}

func (q *MaxQueue)  Push(x interface{})  {
	i := x.(int)
	q.d = append(q.d, i)
}

func (q *MaxQueue) Pop() interface{} {
	var l int
	l, q.d = q.d[len(q.d)-1], q.d[:len(q.d)-1]
	return l
}

func (q *MaxQueue) Len() int {
	return len(q.d)
}


func (q *MaxQueue)  Less(i, j int) bool {
	if q.d[i] > q.d[j] {
		return true
	}
	return false
}


func (q *MaxQueue) Swap(i, j int) {
	tmp := q.d[i]
	q.d[i] = q.d[j]
	q.d[j] = tmp
}