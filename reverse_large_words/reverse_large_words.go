package main

import (
	"fmt"
	"strings"
)

func spinWords(str string) string {
	words := strings.Split(str, " ")

	for idx, word := range words {
		if len(word) >= 5 {
			words[idx] = reverseString(word)
		}
	}

	return strings.Join(words, " ")
}

func reverseString(str string) string {
	runes := []rune(str)
	reversed := ""

	for idx := len(runes) - 1; idx >= 0; idx-- {
		reversed += string(runes[idx])
	}

	return reversed
}

func main() {
	fmt.Println(spinWords("Hey fellow warriors"))  // => returns "Hey wollef sroirraw"
	fmt.Println(spinWords("This is a test"))       // => returns "This is a test"
	fmt.Println(spinWords("This is another test")) // => returns "This is rehtona test"
}
