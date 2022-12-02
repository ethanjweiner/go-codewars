package createphonenumber

import "fmt"

// Complex
func CreatePhoneNumber(numbers [10]uint) string {
	parts := [][]uint{numbers[:3], numbers[3:6], numbers[6:]}
	strParts := make([]string, len(parts))

	for idx, part := range parts {
		strParts[idx] = toString(part)
	}

	return fmt.Sprintf("(%s) %s-%s", strParts[0], strParts[1], strParts[2])
}

func toString(numbers []uint) string {
	str := ""

	for _, number := range numbers {
		str += fmt.Sprint(number)
	}

	return str
}

// Simple
// func CreatePhoneNumber(n [10]uint) string {
// 	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
// }
