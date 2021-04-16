package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// 		nums = []int{1,1,1,2,2,3}, k = 2
		{name: "[]int{1,1,1,2,2,3}", args: args{[]int{1,1,1,2,2,3}, 2}, want: []int{1,2} },
		//		nums = []int{1}, k = 1
		{name: "[]int{1}", args: args{[]int{1}, 1}, want: []int{1} },
		//"[]int{3,0,1,0}, 1
		{name: "[]int{3,0,1,0}", args: args{[]int{1,3,1,2,1}, 1}, want: []int{1} },


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}