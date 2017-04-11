//
package main

import (
	"fmt"
	"time"
)

func main() {
	beginTime := time.Now()
	for i := 0; i < 100; i++ {
		fmt.Println(i, F(i))
	}
	fmt.Println("with F, 耗时", time.Since(beginTime))
}

func F(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return F(n-1) + F(n-2)
}
