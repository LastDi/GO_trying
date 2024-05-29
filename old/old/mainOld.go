package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main3() {
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

func quickSort(arr []int, left int, right int) {
	if left < right {
		partitionIndex := partition(arr, left, right)
		quickSort(arr, left, partitionIndex-1)
		quickSort(arr, partitionIndex+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main1() {
	var maps map[int]int
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

func main2() {
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
