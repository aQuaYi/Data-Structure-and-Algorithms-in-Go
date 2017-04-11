package main

import (
	"fmt"
)

func main() {
	fmt.Println("=====mystery=====")
	fmt.Println("mystery(2,25) =", mystery(2, 25))
	fmt.Println("mystery(3,11) =", mystery(3, 11))
	fmt.Println("=====mystery2=====")
	fmt.Println("mystery2(2,25) =", mystery2(2, 25))
	fmt.Println("mystery2(3,11) =", mystery2(3, 11))

}

func mystery(a, b int) int {
	if b == 0 {
		return 0
	}
	if b%2 == 0 {
		return mystery(a+a, b/2)
	}
	return mystery(a+a, b/2) + a
}

func mystery2(a, b int) int {
	if b == 0 {
		return 1
	}
	if b%2 == 0 {
		return mystery2(a+a, b/2)
	}
	return mystery2(a+a, b/2) * a //题目没有说清楚替换哪个+，所以，我没有替换a+a当中的+
}
