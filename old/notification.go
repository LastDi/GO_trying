package main

import (
	"bufio"
	"fmt"
	"os"
)

var maps map[int]int

func main4() {
	maps = make(map[int]int)

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	count := 0
	global := count
	for i := 0; i < m; i++ {
		var t, id int
		fmt.Fscan(in, &t, &id)
		switch t {
		case 1:
			count++
			switch id {
			case 0:
				global = count
			default:
				maps[id] = count
			}
		case 2:
			if maps[id] < global {
				fmt.Println(global)
			} else {
				fmt.Println(maps[id])
			}
		}
	}
}
