package main

import (
	"fmt"
)

func main() {
	fmt.Println("程序开始")
	s := "ABCdef"
	fmt.Println(s)
	fmt.Println(mystery(s))
}

func mystery(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	a := s[0 : n/2]
	b := s[n/2 : n]
	return b + a
}
