package main

import "fmt"

func digitalRoot(integer int) int {
	if integer/10 < 1 {
		return integer
	}

	return digitalRoot(digitSum(integer))
}

func digitSum(integer int) int {
	if integer >= 0 && integer < 10 {
		return integer
	}

	return digitSum(integer/10) + digitSum(integer%10)
}

func main() {
	fmt.Println(digitalRoot(6))      // 6
	fmt.Println(digitalRoot(10))     // 1
	fmt.Println(digitalRoot(16))     // 7
	fmt.Println(digitalRoot(942))    // 6
	fmt.Println(digitalRoot(132189)) // 6
	fmt.Println(digitalRoot(493193)) // 2
}

/*
16  -->  1 + 6 = 7
942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6
493193  -->  4 + 9 + 3 + 1 + 9 + 3 = 29  -->  2 + 9 = 11  -->  1 + 1 = 2
*/
