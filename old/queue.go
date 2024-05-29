package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var X rune = 88
var Y rune = 89
var Z rune = 90

type Node struct {
	data int
	next *Node
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
		var n int
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		res := find(&s)
		fmt.Println(res)
	}
}

func find(str *string) string {
	arr := []rune(*str)
	var x *Node
	var bx *Node
	var y *Node
	var by *Node
	xc := 0
	yc := 0
	az := strings.Count(*str, "Z")
	ay := strings.Count(*str, "Y")
	ax := strings.Count(*str, "X")
	fmt.Println("Z = ", az, ay, ax)
	for i, ch := range *str {
		switch ch {
		case X:
			if bx == nil {
				x = &Node{data: i, next: nil}
				bx = x
			} else {
				tmp := Node{data: i, next: nil}
				x.next = &tmp
				x = &tmp
			}
			xc++
		case Y:
			if by == nil {
				y = &Node{data: i, next: nil}
				by = y
			} else {
				tmp := Node{data: i, next: nil}
				y.next = &tmp
				y = &tmp
			}
			yc++
		case Z:
			if xc > 1 && yc == 1 {
				if bx.data < by.data {
					arr[bx.data] = -1
					arr[by.data] = -1
					bx = bx.next
					by = by.next
					xc--
					yc--
				}
				arr[i] = -1
				arr[bx.data] = -1
				bx = bx.next
				xc--
			} else if yc > 0 {
				arr[i] = -1
				arr[by.data] = -1
				by = by.next
				yc--
			} else if xc > 0 {
				arr[i] = -1
				arr[bx.data] = -1
				bx = bx.next
				xc--
			} else {
				fmt.Println(bx, by)
				fmt.Println(xc, yc)
				return "No-"
			}
		}
		fmt.Println(arr)
	}

	for bx != nil && by != nil {
		if bx.data >= by.data {
			return "No="
		}
		bx = bx.next
		by = by.next
	}
	if bx != nil || by != nil {
		//fmt.Println(bx, by)
		//fmt.Println(xc, yc)
		return "No+"
	}
	return "Yes"
}
