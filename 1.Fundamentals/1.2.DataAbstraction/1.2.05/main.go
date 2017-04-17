package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("程序开始")
	s := "Hello World"
	strings.ToUpper(s)
	fmt.Println(s[6:11])
	fmt.Println(s)
	fmt.Println("程序结束")
}
