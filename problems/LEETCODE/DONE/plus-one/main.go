package main

import "fmt"

func main() {

	fmt.Println(plusOne([]int{1, 0, 0, 0, 0}))
}

func plusOne(digits []int) []int {
	cur := 0
	rec := 1

	for i := len(digits) - 1; i > -1; i-- {

		cur = digits[i] + rec

		if i == 0 && cur == 10 {
			digits[i] = cur % 10
			res := make([]int, 0)
			res = append(res, 1)
			res = append(res, digits...)
			return res
		}
		if cur%10 == 0 && digits[i] != 0 {
			digits[i] = cur % 10
			rec = 1
			continue
		}

		digits[i] = cur % 10
		rec = 0

	}
	return digits

}
