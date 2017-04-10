package main

import (
	"fmt"
	"math"
)

func main() {
	t := 9.0
	for math.Abs(t-9.0/t) > .001 {
		t = (9.0/t + t) / 2
	}
	fmt.Printf("a.\n%.5f\n", t)
	// 以上是在求t的平方根

	sum := 0
	for i := 1; i < 1000; i++ {
		for j := 0; j < i; j++ {
			sum++
		}
	}
	fmt.Printf("b.\n%d\n", sum)

	sum = 0
	for i := 1; i < 1000; i *= 2 {
		for j := 0; j < 1000; j++ {
			sum++
		}
	}
	fmt.Printf("c.\n%d\n", sum)
}
