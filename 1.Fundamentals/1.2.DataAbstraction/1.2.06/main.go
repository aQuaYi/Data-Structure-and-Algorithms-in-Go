package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("程序开始")
	strA := "ACTGACG"
	strB := "TGACGAC"
	if len(strA) != len(strB) {
		fmt.Println(strA, strB, "不是回文变换")
	}
	if strings.Contains(strA+strA, strB) {
		fmt.Println(strA, strB, "是回文变换")
		return
	}
	fmt.Println(strA, strB, "不是回文变换")
}
