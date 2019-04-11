package main

import (
	"Algorithms-in-Golang/stack"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("程序开始")

	s := "( ( 1 + 2 ) * ( ( 3 - 4 ) * ( 5 - 6 ) ) )"
	fmt.Println(s)
	r := infixToPostfix(s)
	fmt.Println(r)

}

func infixToPostfix(s string) string {
	strs := stack.New()
	ops := stack.New()

	ss := strings.Split(s, " ")
	for _, v := range ss {
		switch v {
		case "(":
			continue
		case "+", "-", "*", "/":
			ops.Push(v)
		case ")":
			b := strs.Pop().(string)
			a := strs.Pop().(string)
			op := ops.Pop().(string)
			strs.Push(fmt.Sprintf("( %s %s %s )", a, b, op))
		default:
			strs.Push(v)
		}
	}
	result := strs.Pop().(string)
	return result
}
