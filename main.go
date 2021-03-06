package main

import (
	"fmt"

	problem1 "github.com/aabishkaryal/Project-Euler/Problem1"
	problem10 "github.com/aabishkaryal/Project-Euler/Problem10"
	problem2 "github.com/aabishkaryal/Project-Euler/Problem2"
	problem3 "github.com/aabishkaryal/Project-Euler/Problem3"
	problem4 "github.com/aabishkaryal/Project-Euler/Problem4"
	problem5 "github.com/aabishkaryal/Project-Euler/Problem5"
	problem6 "github.com/aabishkaryal/Project-Euler/Problem6"
	problem7 "github.com/aabishkaryal/Project-Euler/Problem7"
	problem8 "github.com/aabishkaryal/Project-Euler/Problem8"
	problem9 "github.com/aabishkaryal/Project-Euler/Problem9"
)

func main() {
	fmt.Println("Hello Project Euler")
	fmt.Println("Problem 1: ", problem1.Solution(1000))
	fmt.Println("Problem 2: ", problem2.Solution(4_000_000))
	fmt.Println("Problem 3: ", problem3.Solution(600_851_475_143))
	fmt.Println("Problem 4: ", problem4.Solution(3))
	fmt.Println("Problem 5: ", problem5.Solution(20))
	fmt.Println("Problem 6: ", problem6.Solution(100))
	fmt.Println("Problem 7: ", problem7.Solution(10001))
	fmt.Println("Problem 8: ", problem8.Solution(13))
	fmt.Println("Problem 9: ", problem9.Solution(1000))
	fmt.Println("Problem 10: ", problem10.Solution(2_000_000))
}
