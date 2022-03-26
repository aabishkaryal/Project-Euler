package problem3

func Solution(num int) int {
	largestPrimeFactor, counter := 0, 2
	for counter*counter <= num {
		if num%counter == 0 {
			num /= counter
			largestPrimeFactor = counter
		} else {
			counter++
		}
	}
	if num > largestPrimeFactor {
		largestPrimeFactor = num
	}
	return largestPrimeFactor
}
