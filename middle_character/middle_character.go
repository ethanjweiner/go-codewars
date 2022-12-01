package main

import "fmt"

// func getMiddle(str string) string {
// 	chars := []rune(str)
// 	middleIndex := len(chars) / 2

// 	if len(chars)%2 == 1 {
// 		return fmt.Sprintf("%c", chars[middleIndex])
// 	}

// 	return fmt.Sprintf("%c%c", chars[middleIndex-1], chars[middleIndex])
// }

func getMiddle(str string) string {
	middleIndex := len(str) / 2

	if len(str)%2 == 1 {
		return fmt.Sprintf("%c", str[middleIndex])
	}

	return fmt.Sprintf("%c%c", str[middleIndex-1], str[middleIndex])
}

// Alternative: Return a slice of the string (remember, a string is just a []byte)

func main() {
	fmt.Println(getMiddle("test"))    // should return "es"
	fmt.Println(getMiddle("testing")) // should return "t"
	fmt.Println(getMiddle("middle"))  // should return "dd"
	fmt.Println(getMiddle("A"))       // should return "A"
}
