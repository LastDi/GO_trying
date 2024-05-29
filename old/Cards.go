package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Tup struct {
	index int
	value int
}

func main2() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	var arr []Tup
	arr = make([]Tup, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscan(in, &tmp)
		arr[i] = Tup{i, tmp}
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i].value > arr[j].value })

	var res []int
	res = make([]int, n)

	for _, tup := range arr {
		if tup.value >= m || m < 0 {
			fmt.Println(-1)
			return
		} else {
			res[tup.index] = m
		}
		m--
	}

	for _, el := range res {
		fmt.Print(el, " ")
	}
}
