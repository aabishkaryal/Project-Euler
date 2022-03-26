package main

import (
	"fmt"

	problem1 "github.com/aabishkaryal/Project-Euler/Problem1"
	problem2 "github.com/aabishkaryal/Project-Euler/Problem2"
	problem3 "github.com/aabishkaryal/Project-Euler/Problem3"
)

func main() {
	fmt.Println("Hello Project Euler")
	fmt.Println("Problem 1: ", problem1.Solution(1000))
	fmt.Println("Problem 2: ", problem2.Solution(4_000_000))
	fmt.Println("Problem 3: ", problem3.Solution(600_851_475_143))
}
