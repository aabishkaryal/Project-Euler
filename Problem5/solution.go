package problem5

func Solution(limit int) int {
	lc := 1
	for i := 1; i <= limit; i++ {
		lc = lcm(lc, i)
	}
	return lc
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
