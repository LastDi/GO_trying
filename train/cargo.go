package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Cargo struct {
	ind      int
	start    int
	end      int
	capacity int
}

type Val struct {
	ind int
	val int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for z := 0; z < t; z++ {
		var n int
		fmt.Fscan(in, &n)
		arrival := make([]Val, n)
		for k := 0; k < n; k++ {
			v := Val{}
			fmt.Fscan(in, &v.val)
			v.ind = k
			arrival[k] = v
		}
		var m int
		fmt.Fscan(in, &m)
		cargo := make([]Cargo, m)
		for k := 0; k < m; k++ {
			c := Cargo{}
			fmt.Fscan(in, &c.start, &c.end, &c.capacity)
			c.ind = k
			cargo[k] = c
		}

		sort.Slice(arrival, func(i, j int) bool { return arrival[i].val < arrival[j].val })
		sort.Slice(cargo, func(i, j int) bool {
			if cargo[i].start == cargo[j].start {
				return cargo[i].ind < cargo[j].ind
			} else {
				return cargo[i].start < cargo[j].start
			}
		})

		res := make([]int, n)
		for _, val := range cargo {
			i := 0
			count := 0
			for count < val.capacity && i < len(arrival) {
				if arrival[i].val != -1 {
					if arrival[i].val >= val.start && arrival[i].val <= val.end {
						count++
						arrival[i].val = -1
						res[arrival[i].ind] = val.ind + 1
					}
				}
				i++
			}
		}

		for i := 0; i < len(res); i++ {
			if res[i] == 0 {
				fmt.Print(-1, " ")
			} else {
				fmt.Print(res[i], " ")
			}
		}
		fmt.Println()
	}
}
