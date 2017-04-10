package main

import (
	"fmt"
)

func main() {
	x, y := 0.01, 0.99
	fmt.Println(strictlyCheck(x) && strictlyCheck(y))
}

func strictlyCheck(f float64) bool {
	if 0 < f && f < 1 {
		return true
	}
	return false
}
