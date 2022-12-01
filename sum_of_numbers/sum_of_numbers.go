package main

import "fmt"

// Retrieves sum of sequence of integers [a, b]
func GetSum(a int, b int) int {
	var lesser, more, sum int

	if a < b {
		lesser, more = a, b
	} else {
		lesser, more = b, a
	}

	for current := lesser; current <= more; current++ {
		sum += current
	}

	return sum
}

// Mathematical: Sum = a + (a + 1) + ... + (a + b) / 2 +...+ (b - 1) + b = (b - a + 1) * (a + b) / 21

func main() {
	fmt.Println(GetSum(1, 3))   // 6
	fmt.Println(GetSum(-5, -1)) // -15
	fmt.Println(GetSum(0, 1))   // 1
	fmt.Println(GetSum(1, 1))   // 1
}
