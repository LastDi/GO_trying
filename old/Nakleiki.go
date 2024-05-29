package main

import (
	"bufio"
	"fmt"
	"os"
)

func main3() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str string
	fmt.Fscan(in, &str)
	rs := []rune(str)

	var t int
	fmt.Fscan(in, &t)

	i := 0
	for i < t {
		var a, b int
		var sub string
		fmt.Fscan(in, &a, &b, &sub)

		rsub := []rune(sub)
		j := 0
		for a <= b {
			rs[a-1] = rsub[j]
			j++
			a++
		}
		i++
	}
	fmt.Println(string(rs))
}
