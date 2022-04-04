package problem9

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		sum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{sum: 12}, want: 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.sum); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
