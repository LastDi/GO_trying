package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main5() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for k := 0; k < n; k++ {
		var m int
		fmt.Fscan(in, &m)
		in.ReadLine()

		var jsonByte []byte
		for i := 0; i < m; i++ {
			line, _, _ := in.ReadLine()
			jsonByte = append(jsonByte, line...)
		}

		var jsonn interface{}
		if err := json.Unmarshal(jsonByte, &jsonn); err != nil {
			log.Fatal(err)
		}
		checkInterface(&jsonn)
		//fmt.Println(jsonn)
		marshal, _ := json.Marshal(&jsonn)
		fmt.Println(string(marshal))
	}

}

func checkInterface(val *interface{}) bool {
	switch casted := (*val).(type) {
	case string:
		return false
	case float64:
		return false
	case int:
		return false
	case []interface{}:
		if len(casted) == 0 {
			return true
		} else {
			var del []int
			//del := make([]int, len(casted))
			for i, v := range casted {
				b := checkInterface(&v)
				if b {
					del = append(del, i)
				}
			}
			for i := len(del) - 1; i >= 0; i-- {
				casted = append(casted[:del[i]], casted[del[i]+1:]...)
			}
			*val = casted // вот сука
		}
		if len(casted) == 0 {
			return true
		}
		return false
	case map[string]interface{}:
		if len(casted) == 0 {
			return true
		} else {
			for k, v := range casted {
				b := checkInterface(&v)
				if b {
					delete(casted, k)
				}
			}
		}
		if len(casted) == 0 {
			return true
		}
		return false
	default:
		return false
	}
}
