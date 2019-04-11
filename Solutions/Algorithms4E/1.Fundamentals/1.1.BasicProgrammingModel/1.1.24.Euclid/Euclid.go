package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("请在程序末尾带上两个整数")
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil || a <= 0 {
		log.Fatalln("第一个参数不是一个正确的正整数。")
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil || b <= 0 {
		log.Fatalln("第二个参数不是一个正确的正整数。")
	}
	fmt.Println("最大公约数是", gcd(a, b))
}

func gcd(a, b int) (result int) {
	if a < b {
		a, b = b, a
	}
	fmt.Printf("a=%d, b=%d\n", a, b)

	if b == 0 {
		result = a
		return
	}

	return gcd(b, a%b)
}
