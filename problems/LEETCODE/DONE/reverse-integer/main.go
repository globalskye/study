package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(123 % 10)
	//fmt.Println(123 / 10)
	fmt.Println(reverse(102))
}
func reverse(x int) int {
	res := 0
	cur := 0
	buf := x

	if x < 10 && x > -10 {
		return x
	}
	for {
		if x%10 == 0 {
			x = x / 10
			continue
		}
		break
	}

	for buf >= 9 && x > 10 {
		res = res * 10
		cur = buf % 10
		buf = buf / 10
		res = res + cur
		if buf < 10 && buf > 0 {
			res = res*10 + buf

			if res > math.MaxInt32 {
				return 0
			}
			return res
		}

	}
	for buf <= -9 && x < -10 {
		res = res * 10
		cur = buf % 10
		buf = buf / 10
		res = res + cur
		if buf < 0 && buf > -10 {
			res = res*10 + buf
			if res < math.MaxInt32*-1-1 {
				return 0
			}
			return res
		}

	}

	return x
}
