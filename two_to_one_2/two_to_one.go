package main

import (
	"fmt"
	"sort"
	"strings"
)

func longest(str1 string, str2 string) string {
	// Best practice: allocate appropriate memory for slice (avoids dynamic)
	chars := strings.Split(str1+str2, "")
	sort.Strings(chars)
	return uniqueString(chars)
}

func uniqueString(chars []string) string {
	result := ""

	for _, char := range chars {
		if !strings.Contains(result, char) {
			result += char
		}
	}

	return result
}

func main() {
	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"
	fmt.Println(longest(a, b)) //-> "abcdefklmopqwxy"

	a = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(longest(a, a)) //-> "abcdefghijklmnopqrstuvwxyz"
}
