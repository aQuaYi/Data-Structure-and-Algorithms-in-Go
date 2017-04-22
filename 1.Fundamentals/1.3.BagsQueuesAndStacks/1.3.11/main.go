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
	fs := stack.New()
	ops := stack.New()

	ss := strings.Split(s, " ")
	for _, v := range ss {
		switch v {
		case "(":
			continue
		case "+", "-", "*", "/":
			ops.Push(v)
		case ")":
			b := fs.Pop().(float64)
			a := fs.Pop().(float64)
			op := ops.Pop().(string)
			switch op {
			case "+":
				fs.Push(a + b)
			case "-":
				fs.Push(a - b)
			case "*":
				fs.Push(a * b)
			case "/":
				fs.Push(a / b)
			default:
				fmt.Println("错误的运算符")
				os.Exit(1)
			}
		default:
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Printf("无法把%s转换成float64\n", v)
				os.Exit(1)
			}
			fs.Push(f)
		}
	}
	result := fs.Pop().(float64)
	return result
}
