package main

import "os"
import "log"
import "strconv"
import "fmt"

func main() {
	if len(os.Args) != 4 {
		log.Fatalln("程序尾部添加的整数个数不为3")
	}

	a1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("输入的第一个参数不是整数:", err)
	}

	a2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln("输入的第二个参数不是整数:", err)
	}

	a3, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalln("输入的第三个参数不是整数:", err)
	}

	if a1 != a2 {
		fmt.Println("not equal")
		return
	}

	if a2 != a3 {
		fmt.Println("not equal")
		return
	}

	fmt.Println("equal")
}
