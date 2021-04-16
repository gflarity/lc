package main

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMaxQueue(t *testing.T) {
	q := NewMaxQueue([]int{1,2,3,4,5})
	heap.Init(q)
	assert.EqualValues(t, 5, heap.Pop(q).(int))
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// nums = []int{3,2,1,5,6,4},  k = 2 -> 5
		{"[]int{3,2,1,5,6,4}", args{[]int{3,2,1,5,6,4}, 2}, 5 },
		// nums = []int{3,2,3,1,2,4,5,5,6}, k = 4 -> 4
		{"[]int{3,2,3,1,2,4,5,5,6}", args{[]int{3,2,3,1,2,4,5,5,6}, 4} , 4 },
		// nums = []int{1}, k = 1 -> 1
		{"[]int{1}", args{[]int{1}, 1}, 1 },

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}