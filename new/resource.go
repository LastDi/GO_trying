package main

import (
	"bufio"
	"fmt"
	"os"
)

type Res struct {
	x int
	y int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for z := 0; z < t; z++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		var k int
		fmt.Fscan(in, &k)

		for c := 0; c < k; c++ {

		}
	}
}
