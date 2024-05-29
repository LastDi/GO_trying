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

	var n, t int
	fmt.Fscanln(in, &n, &t)

	mmap := make(map[string]int)

	for i := 0; i < n; i++ {
		var str string
		fmt.Fscan(in, &str)
		if _, ok := mmap[str]; ok {
			mmap[str] = mmap[str] + 1
		} else {
			mmap[str] = 1
		}
	}

	for i := 0; i < t; i++ {
		checked := make(map[string]int)
		var str string
		fmt.Fscan(in, &str)
		for _, val := range str {
			if _, ok := checked[string(val)]; ok {
				checked[string(val)] = checked[string(val)] + 1
			} else {
				checked[string(val)] = 1
			}
		}

		if fmt.Sprint(checked) == fmt.Sprint(mmap) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
