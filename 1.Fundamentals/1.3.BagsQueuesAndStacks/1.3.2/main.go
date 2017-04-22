package main

import (
	"Algorithms-in-Golang/stack"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("程序开始")

	s := "it was - the best - of times - - - it was - the - -"
	ss := strings.Split(s, " ")
	stk := stack.New()
	for _, v := range ss {
		if v != "-" {
			stk.Push(v)
		} else {
			fmt.Print(stk.Pop().(string), " ")
		}
	}
}
