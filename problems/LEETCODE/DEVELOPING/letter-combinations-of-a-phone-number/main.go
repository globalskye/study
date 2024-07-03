package main

import (
	"fmt"
)

func main() {
	fmt.Println('a', 'z')
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	m := make(map[string]string)
	m["2"] = "abc"
	m["3"] = "def"
	m["4"] = "ghi"
	m["5"] = "jkl"
	m["6"] = "mno"
	m["7"] = "pqrs"
	m["8"] = "tuv"
	m["9"] = "wxyz"
	res := make([]string, 0)

	for i := 0; i < len(digits)-1; i++ {
		val  := m[string(digits[i+1])]
		str := ""
		for _, v := range val {
			str +=
		}
		for j := 1; j < len(digits); j++ {

		}

	}

	return res
}
