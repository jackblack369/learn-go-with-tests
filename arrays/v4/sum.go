package main

import "fmt"

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll calculates the respective sums of every slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	// SumAll([]int{1, 2}, []int{0, 9})
	lengthOfNumbers := len(numbersToSum)
	fmt.Println("lengthOfNumbers:", lengthOfNumbers)

	sums := make([]int, lengthOfNumbers)
	fmt.Println("sums:", sums)

	for i, numbers := range numbersToSum {
		fmt.Println("i:", i, "numbers:", numbers)
		sums[i] = Sum(numbers)
		fmt.Println("sums[", i, "]:", sums[i])
	}

	return sums
}
