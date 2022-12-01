package main

import (
	"fmt"
	"sort"
	"strings"
)

func longest(str1 string, str2 string) string {
	// Best practice: allocate appropriate memory for slice (avoids dynamic)
	distinctChars := make([]string, len(str1)+len(str2))

	insertUniqueChars(&distinctChars, str1)
	insertUniqueChars(&distinctChars, str2)

	sort.Strings(distinctChars)

	return strings.Join(distinctChars, "")
}

func insertUniqueChars(chars *[]string, str string) {
	for _, char := range str {
		if !includes(*chars, string(char)) {
			*chars = append(*chars, string(char))
		}
	}
}

func includes[T comparable](slice []T, elementToFind T) bool {
	for _, element := range slice {
		if element == elementToFind {
			return true
		}
	}

	return false
}

func main() {
	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"
	fmt.Println(longest(a, b)) //-> "abcdefklmopqwxy"

	a = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(longest(a, a)) //-> "abcdefghijklmnopqrstuvwxyz"
}
