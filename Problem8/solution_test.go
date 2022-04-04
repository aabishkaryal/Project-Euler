package problem8

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		numDigits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{numDigits: 4}, want: 5832},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.numDigits); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
