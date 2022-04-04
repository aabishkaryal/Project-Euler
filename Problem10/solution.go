package problem10

func Solution(limit int) int {
	isComposite := make([]bool, limit)
	sum := 0
	for p := 2; p < limit; p++ {
		if isComposite[p] {
			continue
		}

		sum += p
		for i := p * p; i < limit; i += p {
			isComposite[i] = true
		}
	}
	return sum
}
