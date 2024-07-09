package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nthUglyNumber(11))
}
func nthUglyNumber(n int) int {
	arr := make([]int, 0)
	val := 1
	idx := 0
	arr = append(arr, 1)
	for {
		math.Log(1)
		if val%2 == 0 || val%3 == 0 || val%5 == 0 {
			arr = append(arr, val)
			if len(arr) == n {
				return arr[len(arr)-1]
			}
			idx++
		}
		val++

	}
}
