package problem7

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		nth int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nth: 6}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.nth); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
