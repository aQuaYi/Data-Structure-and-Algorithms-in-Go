package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(exR1(6))
}

func exR1(n int) string {
	if n <= 0 {
		return ""
	}
	return exR1(n-3) + strconv.Itoa(n) + exR1(n-2) + strconv.Itoa(n)
}
