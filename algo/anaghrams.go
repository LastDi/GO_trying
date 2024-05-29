package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	anagrams := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(anagrams)
}

func groupAnagrams(strs []string) [][]string {
	mapp := make(map[string][]string)
	res := [][]string{}
	for _, str := range strs {
		char := strings.Split(str, "")
		sort.Strings(char)
		sorted := strings.Join(char, "")
		mapp[sorted] = append(mapp[sorted], str)
	}

	for _, str := range mapp {
		res = append(res, str)
	}
	return res
}
