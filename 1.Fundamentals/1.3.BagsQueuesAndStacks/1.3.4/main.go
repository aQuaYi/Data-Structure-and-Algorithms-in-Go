package main

import (
	"Algorithms-in-Golang/stack"
	"fmt"
	"strings"
)

var (
	p = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
)

func main() {
	fmt.Println("程序开始")
	check("[(])")
	check("[()]{}{[()()]()}")
}

func check(s string) {
	stk := stack.New()
	ss := strings.Split(s, "")

	for _, v := range ss {
		switch v {
		case "(", "[", "{":
			stk.Push(v)
		case ")", "]", "}":
			if stk.Pop().(string) != p[v] {
				fmt.Println("false")
				return
			}
		default:
			fmt.Printf("%s当中有其他字符\n", s)
		}
	}
	fmt.Println("true")
}
