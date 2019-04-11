//我对题目的理解是
//编写一个函数ln(n)去计算n！的值
package main

import (
	"fmt"
)

func main() {
	fmt.Println("ln(2) =", ln(2))
	fmt.Println("ln(3) =", ln(3))
	fmt.Println("ln(5) =", ln(5))
	fmt.Println("ln(10) =", ln(10))
}

func ln(n int) int {
	if n == 1 {
		return 1
	}
	return n * ln(n-1)
}
