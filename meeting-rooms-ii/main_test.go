package main

import "testing"

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "[][]int{{0,30},{5,10},{15,20}}", args: args{ intervals: [][]int{{0,30},{5,10},{15,20}}}, want: 2},
		{name: "[][]int{{7,10},{2,4}}", args: args{ intervals: [][]int{{7,10},{2,4}}}, want: 1},
		{name: "[][]int{{0,30}, {0,30}, {0,30}}", args: args{ intervals: [][]int{{0,30}, {0,30}, {0,30}}}, want: 3},
		{name: "[][]int{{2,15},{36,45},{9,29},{16,23},{4,9}}", args: args{ intervals: [][]int{{2,15},{36,45},{9,29},{16,23},{4,9}}}, want: 2},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
