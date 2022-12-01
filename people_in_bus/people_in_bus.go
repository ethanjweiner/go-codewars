package main

import "fmt"

func PeopleRemaining(busStops [][2]int) int {
	result := 0

	for _, busStop := range busStops {
		result += busStop[0] - busStop[1]
	}

	return result
}

func main() {
	fmt.Println(PeopleRemaining([][2]int{{10, 0}, {3, 5}, {5, 8}})) // 5
	fmt.Println(PeopleRemaining([][2]int{{0, 0}}))                  // 0
}
