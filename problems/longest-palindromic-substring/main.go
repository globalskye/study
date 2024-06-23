package main

import (
	"fmt"
	"strings"
)

func main() {
	longestPalindrome("dsadasdasaabababab")
	fmt.Println()
}

type T struct {
}

func longestPalindrome(s string) string {

	//var maxLen int
	//var currentLen int
	//var start int
	var substrings []string
	for i := 0; i < len(s); i++ {
		currentValue := s[i]

		res := strings.IndexByte(s[i+1:], currentValue)
		fmt.Println(s[i+1:])
		if res == -1 {
			continue
		}
		if res == 0 {
			res = i
		}
		substr := s[i : res+i]
		for j := 0; j < len(substr)/2; j++ {

			if len(substr)%2 == 0 {
				if substr[len(substr)/2+j] == substr[len(substr)/2-1-j] {
					if j == len(substr)/2-1 {
						substrings = append(substrings, s[i:res+2])
					}
					continue
				} else {
					break
				}
			}
			if len(substr)%2 == 1 {
				if substr[len(substr)/2+j+1] == substr[len(substr)/2-1-j] {
					if j == len(substr)/2-1 {
						substrings = append(substrings, s[i:res+2])
					}
					continue
				} else {
					break
				}
			}

		}
		fmt.Println(substrings)

	}
	return "0"
}
