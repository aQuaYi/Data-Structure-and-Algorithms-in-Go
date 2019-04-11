package main

import (
	"Algorithms-in-Golang/stack"
	"fmt"
)

func main() {
	fmt.Println("程序开始")
	n := 50
	fmt.Printf("%d=", n)
	stk := stack.New()
	for n > 0 {
		stk.Push(n % 2)
		n = n / 2
	}
	c := stk.Iterator()
	for c.HasNext() {
		fmt.Print(c.Next().(int))
	}
	fmt.Println()
}
