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

	var n, m int
	fmt.Fscan(in, &n, &m)

	var arr []int
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	fmt.Println(arr)
	var sorted = make([]int, len(arr))
	copy(sorted, arr)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	fmt.Println(sorted)

	var mapSort map[int]int
	mapSort = make(map[int]int)

	for i := 0; i < len(sorted); i++ {
		if _, ok := mapSort[sorted[i]]; !ok {
			mapSort[sorted[i]] = i
		}
	}
	var qual = make([]int, len(sorted))
	qual[0] = sorted[0]
	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			qual[i] = qual[i-1] + 1
		} else {
			fmt.Println(sorted[i], qual[i-1])
			if sorted[i] <= qual[i-1] {
				fmt.Println(-1)
				return
			}
			qual[i] = sorted[i]
		}
	}

	var res []int
	res = make([]int, n)
	//fmt.Println("before", mapSort)
	for i := 0; i < n; i++ {
		//fmt.Println(arr[i], mapSort[arr[i]], sorted[mapSort[arr[i]]])
		res[i] = sorted[mapSort[arr[i]]] + 1
		if res[i] > m {
			fmt.Println(-1)
			return
		}
		sorted[mapSort[arr[i]]] = res[i]
		//fmt.Println("after", mapSort, sorted)
	}

	fmt.Println(mapSort)
	fmt.Println(arr)
	fmt.Println(res)

	for _, r := range res {
		fmt.Print(r, " ")
	}
}
