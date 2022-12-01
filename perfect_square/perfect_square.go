package main

import (
	"fmt"
	"math"
)

func findNextSquare(currentSquare float64) int {
	squareRoot := math.Sqrt(currentSquare)

	if !isInteger(squareRoot) {
		return -1
	}

	return int(math.Pow(squareRoot+1, 2))
}

func isInteger(num float64) bool {
	return num == math.Trunc(num)
}

func main() {
	fmt.Println(findNextSquare(121)) // 144
	fmt.Println(findNextSquare(625)) // 676
	fmt.Println(findNextSquare(114)) // -1
}
