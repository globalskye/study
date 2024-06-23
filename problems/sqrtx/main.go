package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(10))
	fmt.Println(mySqrt(10))
}
func mySqrt(x int) int {
	i := 0
	res := 0
	for i < x {
		res = x / 2
		tmp := res * res
		if tmp > x {
			res = res / 2
		}
		if tmp < x {
			res = res + res/2
		}
		i++
	}
	return res
}
