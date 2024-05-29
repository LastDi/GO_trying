package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, b, r, f int
		fmt.Fscan(in, &n, &b, &r, &f)

		var arr map[string]int
		arr = make(map[string]int)

		var blueSubs map[string]int
		blueSubs = make(map[string]int)

		for i := 0; i < n; i++ {
			var s string
			fmt.Fscan(in, &s)
			arr[s] = i
		}

		for str, i := range arr {
			//fmt.Println(blueSubs)
			//fmt.Println(str, " --")
			// здесь набиваем мапу с синими подстроками
			if i < b {
				tmp := make(map[string]int)
				for j := 0; j < len(str)-1; j++ {
					for k := j + 1; k <= len(str); k++ {
						sub := str[j:k]
						if _, ok := arr[sub]; !ok {
							if _, okk := tmp[sub]; !okk {
								tmp[sub] = 0
								blueSubs[sub] = blueSubs[sub] + 1
							}
						}
					}
				}
			} else if i < b+r {
				// здесь декрементим счетчик если подстрока из красного слова есть в мапе синих
				tmpt := make(map[string]int)
				for j := 0; j < len(str)-1; j++ {
					for k := j + 1; k <= len(str); k++ {
						sub := str[j:k]
						if _, ok := blueSubs[sub]; ok {
							if _, okk := tmpt[sub]; !okk {
								fmt.Println("==============")
								tmpt[sub] = 0
								blueSubs[sub] = blueSubs[sub] - 1
							}
						}
					}
				}
			} else if i == f-1 {
				// здесь удаляем из мапы подстроку если он есть в составе черного слова
				for j := 0; j < len(str)-1; j++ {
					for k := j + 1; k <= len(str); k++ {
						//sub := str[j:k]
						delete(blueSubs, str[j:k])
						//if _, ok := blueSubs[sub]; ok {
						//}
					}
				}
			}
		}

		max := 0
		maxVal := ""
		for val, key := range blueSubs {
			if key > max {
				max = key
				maxVal = val
			} else if key == max && len(val) > len(maxVal) {
				max = key
				maxVal = val
			}
		}
		if maxVal == "" {
			maxVal = "tkhapjiabb"
		}

		//fmt.Println(blueSubs)
		//fmt.Println()
		fmt.Println(maxVal, max)
	}
}
