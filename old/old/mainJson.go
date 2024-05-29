package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main312() {
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

		//dec := json.NewDecoder(in)
		//for {
		//	var v Folder
		//	if err := dec.Decode(&v); err != nil {
		//		if err == io.EOF {
		//			break
		//		}
		//		log.Fatal(err)
		//	}
		//	fmt.Println(v)
		//}

		var root Folder
		err := json.NewDecoder(in).Decode(&root)
		if err != nil {
			log.Fatal(err)
		}
		var res int
		//recursiveFolderCheck1(&root, false, &res)
		root.recursiveFolderCheck1(false, &res)
		fmt.Println(res)
	}
}

func (folder Folder) recursiveFolderCheck1(inf bool, res *int) {
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
		recursiveFolderCheck1(&fld, inf, res)
	}
}

func recursiveFolderCheck1(folder *Folder, inf bool, res *int) {
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
		recursiveFolderCheck1(&fld, inf, res)
	}
}
