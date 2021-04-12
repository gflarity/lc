package main

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "()())", args: args{ s: "()())"}, want: "()()"},
		{name: "(((b))c)", args: args{ s: "(((b))c)"}, want: "(((b))c)"},
		{name: "))((", args: args{ s: "))(("}, want: ""},
		{name: "a)b(c)d", args: args{ s: "a)b(c)d"}, want: "ab(c)d"},
		{name: "((())", args: args{ s: "((())"}, want: "(())"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}