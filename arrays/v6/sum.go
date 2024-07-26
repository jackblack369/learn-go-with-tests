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

// SumAllTails calculates the respective sums of every slice passed in.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
		fmt.Println("numbers:", numbers, " tail:", tail, " sums:", sums)
	}

	return sums
}
