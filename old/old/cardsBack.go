package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Counter struct {
	count int
	index int
}

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
	//var maps map[int]int
	//maps = make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		//maps[i] = arr[i]
	}

	fmt.Println(arr)
	var sorted = make([]int, len(arr))
	copy(sorted, arr)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	fmt.Println(sorted)

	var mapSort map[int]Counter
	mapSort = make(map[int]Counter)

	//i := sorted[0]
	//mapSort[sorted[0]] = 1
	for i := 0; i < len(sorted); i++ {
		if val, ok := mapSort[sorted[i]]; !ok {
			mapSort[sorted[i]] = Counter{1, i}
		} else {
			mapSort[sorted[i]] = Counter{val.count + 1, i}
		}
	}

	var res []int
	res = make([]int, n)
	fmt.Println("before", mapSort)
	for i := 0; i < n; i++ {
		fmt.Println(arr[i], mapSort[arr[i]], sorted[mapSort[arr[i]].index])
		res[i] = sorted[mapSort[arr[i]].index] + 1

		if res[i] > m {
			fmt.Println(-1)
			return
		}
		//fmt.Println("after", mapSort, sorted)
	}

	fmt.Println(n, m, mapSort)
	fmt.Println(arr)
	fmt.Println(res)
}
