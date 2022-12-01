package main

import (
	"fmt"
	"strings"
)

func accum(str string) string {
	var substrings []string

	for idx, char := range str {
		substrings = append(substrings, strings.ToUpper(string(char))+strings.Repeat(strings.ToLower(string(char)), idx+1))
	}

	return strings.Join(substrings, "-")
}

func main() {
	fmt.Println(accum("abcd"))    // -> "A-Bb-Ccc-Dddd"
	fmt.Println(accum("RqaEzty")) // -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
	fmt.Println(accum("cwAt"))    // -> "C-Ww-Aaa-Tttt"
}
