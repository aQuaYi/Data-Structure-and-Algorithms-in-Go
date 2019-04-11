package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1024
	s := ""
	for n > 0 {
		s = strconv.Itoa(n%2) + s
		n /= 2
	}
	fmt.Println(s)
}
