package problem10

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{limit: 10}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.limit); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

var inputs = []int{100, 1000, 10000, 100000}

func BenchmarkSolution(b *testing.B) {
	for _, input := range inputs {
		b.Run(fmt.Sprintf("Input Size: %v", input), func(b *testing.B) {
			for i := 0; i <= b.N; i++ {
				_ = Solution(input)
			}
		})
	}
}
