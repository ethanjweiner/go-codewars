package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type byNumber []string

func (n byNumber) Len() int {
	return len(n)
}

func (n byNumber) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n byNumber) Less(i, j int) bool {
	iNum, _ := strconv.Atoi(n[i])
	jNum, _ := strconv.Atoi(n[j])
	return iNum < jNum
}

func HighAndLow(numbersString string) string {
	numbers := strings.Split(numbersString, " ")
	sort.Sort(byNumber(numbers))
	return fmt.Sprintf("%s %s", numbers[len(numbers)-1], numbers[0])
}

func main() {
	fmt.Println(HighAndLow("1 2 3 4 5"))  // return "5 1"
	fmt.Println(HighAndLow("1 2 -3 4 5")) // return "5 -3"
	fmt.Println(HighAndLow("1 9 3 4 -5")) // return "9 -5"
}
