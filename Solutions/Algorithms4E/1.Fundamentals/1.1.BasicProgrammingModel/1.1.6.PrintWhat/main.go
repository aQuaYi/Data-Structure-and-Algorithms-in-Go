package main

import "fmt"

func main() {
	f, g := 0, 1
	for i := 0; i <= 15; i++ {
		fmt.Println(f)
		f = f + g
		g = f - g
	}
}
