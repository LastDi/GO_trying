package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main1() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		var res []int
		res = make([]int, t)

		for i := 0; i < t; i++ {
			fmt.Fscan(in, &res[i])
		}

		var copied = make([]int, len(res))
		copy(copied, res)

		sort.Slice(copied, func(i, j int) bool {
			return copied[i] < copied[j]
		})
		//quickSort(copied, 0, len(copied)-1)

		var maps map[int]int
		maps = make(map[int]int)

		for ind, el := range copied {
			if maps[el-1] != 0 {
				maps[el] = maps[el-1]
			} else if maps[el] == 0 {
				maps[el] = ind + 1
			}
		}

		for _, el := range res {
			fmt.Print(maps[el], " ")
		}
		fmt.Println()
	}
}
