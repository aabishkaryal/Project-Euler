package problem4

import (
	"math"
)

func Solution(numDigits int) int {
	var largestPalindromeProduct int
	for i := largestNum(numDigits); i >= smallestNum(numDigits); i-- {
		for j := i; j >= smallestNum(numDigits); j-- {
			if i*j > largestPalindromeProduct && isPalindrome(i*j) {
				largestPalindromeProduct = i * j
			} else if i*j < largestPalindromeProduct {
				break
			}
		}
	}
	return largestPalindromeProduct
}

// Returns whether the given string is palindrome
func isPalindrome(n int) bool {
	original := n
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}

// Returns the largest number with n digits
func largestNum(n int) int {
	return smallestNum(n+1) - 1
}

// Returns the smallest number with n digits
func smallestNum(n int) int {
	return int(math.Pow10(n - 1))
}
