package main

import (
	"Algorithms-in-Golang/stack"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("程序开始")

	strs := stack.New()
	ops := stack.New()

	s := "1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )"
	ss := strings.Split(s, " ")
	for _, v := range ss {
		switch v {
		case "+", "-", "*", "/":
			ops.Push(v)
		case ")":
			b := strs.Pop().(string)
			a := strs.Pop().(string)
			op := ops.Pop().(string)
			strs.Push(fmt.Sprintf("( %s %s %s )", a, op, b))
		default:
			strs.Push(v)
		}
	}
	result := strs.Pop().(string)

	fmt.Println(s)
	fmt.Println(result)
}
