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

func main() {
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
		var root Folder
		err := json.NewDecoder(in).Decode(&root)
		if err != nil {
			log.Fatal(err)
		}
		var res int
		folderCheck(&root, false, &res)
		fmt.Println(res)
	}
}

func folderCheck(folder *Folder, inf bool, res *int) {
	if !inf {
		for _, f := range folder.Files {
			matched, _ := regexp.MatchString(".\\.hack$", f)
			if matched {
				inf = matched
				*res += len(folder.Files)
				break
			}
		}
	} else {
		*res += len(folder.Files)
	}

	for _, fld := range folder.Folders {
		folderCheck(&fld, inf, res)
	}
}
