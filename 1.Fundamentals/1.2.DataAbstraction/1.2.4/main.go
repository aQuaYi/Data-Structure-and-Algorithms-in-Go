package main

import (
	"fmt"
)

func main() {
	fmt.Println("程序开始")
	str1 := "hello"
	str2 := str1
	str1 = "world"
	fmt.Println(str1)
	fmt.Println(str2)
}
