package problem4

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{n: 123}, want: false},
		{name: "test2", args: args{n: 12321}, want: true},
		{name: "test3", args: args{n: 123321}, want: true},
		{name: "test4", args: args{n: 121}, want: true},
		{name: "test5", args: args{n: 11}, want: true},
		{name: "test6", args: args{n: 12}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.n); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smallestNum(t *testing.T) {
	type args struct {
		digits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{digits: 2}, want: 10},
		{name: "test2", args: args{digits: 3}, want: 100},
		{name: "test3", args: args{digits: 4}, want: 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNum(tt.args.digits); got != tt.want {
				t.Errorf("smallestNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestNum(t *testing.T) {
	type args struct {
		digits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{digits: 2}, want: 99},
		{name: "test2", args: args{digits: 3}, want: 999},
		{name: "test3", args: args{digits: 4}, want: 9999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNum(tt.args.digits); got != tt.want {
				t.Errorf("largestNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution(t *testing.T) {
	type args struct {
		numDigits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{numDigits: 2}, want: 9009},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.numDigits); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
