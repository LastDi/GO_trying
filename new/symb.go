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

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var str string
		fmt.Fscan(in, &str)
		res := "YES"
		if len(str) < 2 {
			res = "YES"
		} else {
			ch := str[0]
			ins := true
			for i := 1; i < len(str); i++ {
				if ins {
					if str[i] != ch {
						ins = false
					}
				} else {
					if str[i] != ch {
						res = "NO"
					} else {
						ins = true
					}
				}
			}
			if !ins {
				res = "NO"
			}
		}
		fmt.Println(res)
	}
}
