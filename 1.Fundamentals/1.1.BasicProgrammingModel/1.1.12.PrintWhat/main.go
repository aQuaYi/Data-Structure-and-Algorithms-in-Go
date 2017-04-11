package main

import (
	"fmt"
)

func main() {
	a := [10]int{}
	for i := 0; i < 10; i++ {
		a[i] = 9 - i
	}
	for i := 0; i < 10; i++ {
		a[i] = a[a[i]]
	}
	for i := 0; i < 10; i++ {
		fmt.Println(a[i]) //原题是错的，应该是打印a[i]
	}
}
