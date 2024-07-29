package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var count int

	fmt.Fscan(in, &count)
	var n int
	var p float64

	var res float64
	for i := 0; i < count; i++ {
		fmt.Fscan(in, &n, &p)
		for j := 0; j < n; j++ {
			var a float64
			fmt.Fscan(in, &a)

			commission := a * p / 100

			var c int
			c = int(a*p) / 100

			res += commission - float64(c)

		}
		fmt.Fprintf(out, "%.2f\n", res)
		res = 0
	}

}
