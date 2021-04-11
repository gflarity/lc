package main

import (
	"testing"
)

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ name: "111111111111111111111111111111111111111111111", args: args{s: "111111111111111111111111111111111111111111111"}, want: 1836311903 },
		{ "12", args{ "12" }, 2},

		// 2 26, 2 2 6, 22, 6
		//
		{ "226", args{ "226" }, 3},
		{ "416", args{ "416" }, 2},
		{ "0", args{ "0" }, 0},
		{ "06", args{ "06" }, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

