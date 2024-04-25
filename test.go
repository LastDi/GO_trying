package main

import (
	"fmt"
	"regexp"
)

func main111() {
	matched, err := regexp.MatchString(".\\.hack$", "s.hack")
	fmt.Println(matched, err)

}
