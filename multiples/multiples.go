package main

import "fmt"

func sumOfMultiples(upperBound int) int {
	sum := 0

	for num := 3; num < upperBound; num++ {
		if num%5 == 0 || num%3 == 0 {
			sum += num
		}
	}

	return sum
}

func main() {
	fmt.Println(sumOfMultiples(-5)) // 0
	fmt.Println(sumOfMultiples(0))  // 0
	fmt.Println(sumOfMultiples(2))  // 0
	fmt.Println(sumOfMultiples(3))  // 0
	fmt.Println(sumOfMultiples(5))  // 3
	fmt.Println(sumOfMultiples(10)) // 23
	fmt.Println(sumOfMultiples(13)) // 45
	fmt.Println(sumOfMultiples(16)) // 60
}
