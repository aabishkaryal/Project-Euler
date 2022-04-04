package problem9

func Solution(sum int) int {
	for m := 1; m <= sum; m++ {
		for n := 1; n <= m; n++ {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			if a+b+c == sum {
				return a * b * c
			}
		}
	}
	return 0
}
