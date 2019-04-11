package main

import (
	"fmt"
	"log"
)

func main() {
	N := 1<<63 - 1
	fmt.Println(N, lg(N))
}

func lg(n int) int {
	if n <= 0 {
		log.Fatalln("lg的参数必须为正数")
	}
	exp := 0
	base := 2.0
	for base <= float64(n) { //把n转换为float64是为了可以把int64最大的数都检测出来。
		base *= 2
		exp++
	}
	return exp
}
