package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Tuple struct {
	index int
	value int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		var arr []Tuple
		arr = make([]Tuple, m)
		var res []string
		res = make([]string, m)

		for i := 0; i < m; i++ {
			var r int
			fmt.Fscan(in, &r)
			arr[i] = Tuple{i, r}
			res[i] = "0"
		}
		sort.Slice(arr, func(i, j int) bool { return arr[i].value < arr[j].value })

		flag := true
		if arr[0].value != 1 {
			arr[0].value = arr[0].value - 1
			res[arr[0].index] = "-"
		}
		for i := 1; i < m; i++ {
			if arr[i].value == arr[i-1].value {
				if arr[i].value < n {
					arr[i].value = arr[i].value + 1
					res[arr[i].index] = "+"
				} else {
					fmt.Print("x")
					flag = false
					break
				}
			} else if arr[i].value < arr[i-1].value {
				fmt.Print("x")
				flag = false
				break
			} else {
				if i+1 < m && (arr[i].value-arr[i-1].value) > 1 {
					arr[i].value = arr[i].value - 1
					res[arr[i].index] = "-"
				}
			}
		}

		if flag {
			for _, el := range res {
				fmt.Print(el)
			}
		}
		fmt.Println()
	}
}
