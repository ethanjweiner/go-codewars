package main

import (
	"fmt"
)

func VowelCount(str string) int {
	VOWELS := []rune{'a', 'e', 'i', 'o', 'u'}

	count := 0

	for _, char := range str {
		if includes(VOWELS, char) {
			count += 1
		}
	}

	return count
}

func includes[T comparable](slice []T, valToFind T) bool {
	for _, val := range slice {
		if val == valToFind {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(VowelCount("abcdef")) // 2
}
