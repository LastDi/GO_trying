package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main123() {
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

		var root Folder
		if err := json.Unmarshal(jsonByte, &root); err != nil {
			log.Fatal(err)
		}
		var res int
		recursiveFolderCheck(root, false, &res)
		fmt.Println(res)
	}
}

func recursiveFolderCheck(folder Folder, inf bool, res *int) {
	for _, f := range folder.Files {
		matched, _ := regexp.MatchString(".\\.hack$", f)
		inf = inf || matched
		if inf {
			*res += len(folder.Files)
			break
		}
	}
	for _, fld := range folder.Folders {
		recursiveFolderCheck(fld, inf, res)
	}
}
