package main

import "fmt"

func evenOrOdd(value int) string {
	if value%2 == 0 {
		return "even"
	}

	return "odd"
}

func main() {
	fmt.Println(evenOrOdd(0))
	fmt.Println(evenOrOdd(2))
	fmt.Println(evenOrOdd(3))
}
