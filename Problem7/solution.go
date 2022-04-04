package problem7

func Solution(nth int) int {
	primes := make([]int, 0, nth)
	n := 2
	for len(primes) < nth {
		isPrime := true
		for _, p := range primes {
			if n%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, n)
		}
		n++
	}
	return primes[len(primes)-1]
}
