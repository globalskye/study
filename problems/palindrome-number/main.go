package main

import (
	"fmt"
)

func main() {
	fmt.Print(isPalindrome(554212455))
}
func isPalindrome(x int) bool {
	var reverse int
	current := x
	if x < 0 {
		return false
	}
	for {
		remainder := x % 10
		reverse = reverse*10 + remainder
		x = x / 10

		if x == 0 {
			return current == reverse
		}
	}

	return false
}
