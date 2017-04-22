package main

import (
	"Algorithms-in-Golang/stack"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("程序开始")

	s := "( ( 1 2 + ) ( ( 3 4 * ) ( 5 7 / ) * ) * )"
	fmt.Println(s)
	r := evaluatePostfix(s)
	fmt.Println(r)

}

func evaluatePostfix(s string) float64 {
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
			b, _ := strconv.ParseFloat(strs.Pop().(string), 64)
			a, _ := strconv.ParseFloat(strs.Pop().(string), 64)
			op := ops.Pop().(string)
			switch op {
			case "+":
				strs.Push(fmt.Sprint(a + b))
			case "-":
				strs.Push(fmt.Sprint(a - b))
			case "*":
				strs.Push(fmt.Sprint(a * b))
			case "/":
				strs.Push(fmt.Sprint(a / b))
			default:
				fmt.Println("错误的运算符")
				os.Exit(1)
			}
		default:
			strs.Push(v)
		}
	}
	result, _ := strconv.ParseFloat(strs.Pop().(string), 64)
	return result
}
