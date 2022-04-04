package problem6

func Solution(limit int) int {
	return squareOfSum(limit) - sumOfSquares(limit)
}

func sumOfSquares(n int) int {
	return n * (2*n + 1) * (n + 1) / 6
}

func squareOfSum(n int) int {
	return n * n * (n + 1) * (n + 1) / 4
}
