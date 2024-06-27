package main

import (
	"fmt"
)

func main() {
	fmt.Println('A', 'B')
	fmt.Println(string(rune(89)))
	fmt.Println('@')
	fmt.Println(convertToTitle(53))

}
func convertToTitle(columnNumber int) string {
	count := 26
	base := 65
	result := ""
	//28 - AB
	//52 - AZ
	//53 - BA
	//701 - ZY
	for columnNumber > 0 {
		columnNumber = columnNumber - 1
		currentNumber := columnNumber % count
		columnNumber = columnNumber / count

		result = string(rune(base+currentNumber)) + result

	}

	return result
}
