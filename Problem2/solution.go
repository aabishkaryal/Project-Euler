package problem2

func Solution(limit int) int {
	sum := 0
	a, b := 1, 2
	c := a + b
	for b < limit {
		sum += b
		a = b + c
		b = c + a
		c = a + b

	}
	return sum
}
