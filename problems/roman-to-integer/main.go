package main

import "fmt"

func main() {
	fmt.Print(romanToInt("LVIII"))
}
func romanToInt(s string) int {
	var res int
	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			current := string(s[len(s)-i-1])
			currentNumber, _ := m[current]
			res += currentNumber
			continue
		}
		current := string(s[len(s)-i-1])
		next := string(s[len(s)-i-2])
		currentNumber, _ := m[current]
		nextNumber, _ := m[next]

		if currentNumber > nextNumber {
			res += currentNumber - nextNumber
			i = i + 1
			continue
		}
		res += currentNumber

	}
	return res
}
