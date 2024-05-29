package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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
		var q int
		fmt.Fscan(in, &q)
		arr := make([]int, q)
		for j := 0; j < q; j++ {
			var g int
			fmt.Fscan(in, &g)
			w := 1
			for k := 0; k < g; k++ {
				w *= 2
			}
			arr[j] = w
		}

		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
		fmt.Println(arr)

		i := 0
		res := 0
		var check []int
		for i < len(arr) {
			w := 0
			for w <= m && i < len(arr) {
				if arr[i]+w <= m {
					w += arr[i]
					arr[i] = 0
					i++
				} else {
					k := i
					for k < len(arr) && (arr[k] == 0 || arr[k]+w > m) {
						k++
					}
					if k < len(arr) {
						w += arr[k]
						arr[k] = 0
					} else {
						break
					}
				}
			}
			check = append(check, w)
			res++
		}

		r := res/n + res%n
		fmt.Println(r)
		fmt.Println(check)
	}
}
