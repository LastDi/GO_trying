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

	//var sb strings.Builder

	for i := 0; i < t; i++ {
		var n, b, r, f int
		fmt.Fscan(in, &n, &b, &r, &f)

		var arrMap map[string]int
		arrMap = make(map[string]int)
		var arr []string
		arr = make([]string, n)

		var blueSubs map[string]int
		blueSubs = make(map[string]int)

		for i := 0; i < n; i++ {
			var s string
			fmt.Fscan(in, &s)
			arr[i] = s
			arrMap[s] = i
		}

		for ind, str := range arr {
			// здесь набиваем мапу с синими подстроками
			if ind < b {
				var tmp map[string]int
				tmp = make(map[string]int)
				for j := 0; j < len(str); j++ {
					for k := j + 1; k <= len(str); k++ {
						sub := str[j:k]
						if _, ok := arrMap[sub]; !ok {
							if _, okk := tmp[sub]; !okk {
								tmp[sub] = 0
								blueSubs[sub] = blueSubs[sub] + 1
							}
						}
					}
				}
			} else if ind < b+r {
				// здесь декрементим счетчик если подстрока из красного слова есть в мапе синих
				var tmpRed map[string]int
				tmpRed = make(map[string]int)
				for j := 0; j < len(str); j++ {
					for k := j + 1; k <= len(str); k++ {
						sub := str[j:k]
						if _, ok := blueSubs[sub]; ok {
							if _, okk := tmpRed[sub]; !okk {
								tmpRed[sub] = 0
								blueSubs[sub] = blueSubs[sub] - 1
							}
						}
					}
				}
			} else if ind == f-1 {
				// здесь удаляем из мапы подстроку если он есть в составе черного слова
				for j := 0; j < len(str); j++ {
					for k := j + 1; k <= len(str); k++ {
						delete(blueSubs, str[j:k])
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

		//sb.WriteString(maxVal)
		//sb.WriteString(" ")
		//sb.WriteString(strconv.Itoa(max))
		//sb.WriteString("\n")

		fmt.Println(maxVal, max)
	}

	//fmt.Println(sb.String())
}
