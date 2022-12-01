package main

import "fmt"

// Add error handling
func findOddInt(integers []int) int {
	for integer, count := range tally(integers) {
		if count%2 == 1 {
			return integer
		}
	}

	return 0
}

func tally[T comparable](slice []T) map[T]int {
	counts := make(map[T]int, len(slice)) // Initialized w/ values of `0`

	for _, item := range slice {
		counts[item] += 1
	}

	return counts
}

func main() {
	fmt.Println(findOddInt([]int{7}))             // 7
	fmt.Println(findOddInt([]int{1, 1, 2}))       // 2
	fmt.Println(findOddInt([]int{0, 1, 0, 1, 0})) // 0
}
