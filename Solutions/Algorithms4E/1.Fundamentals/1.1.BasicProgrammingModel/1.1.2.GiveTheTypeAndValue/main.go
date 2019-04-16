package main

import (
	"fmt"
)

func main() {
	a := (1 + 2.236) / 2
	fmt.Printf("%T\t%f\n", a, a)
	b := 1 + 2 + 3 + 4.0
	fmt.Printf("%T\t%f\n", b, b)
	c := 4.1 >= 4
	fmt.Printf("%T\t%t\n", c, c)
	// 取消下面两行的注释符号，执行文件，程序会报错。因为Golang是强类型语言，进行运算的量需要类型一致。
	// d := 1 + 2 + "3"
	// fmt.Printf("%T\t%s\n", c, c)

	//上面a变量中 1+2.236的两个数字类型不一致，却能执行成功的原因是，Golang对常数和constant类型的恒量有一定的宽容性，在计算时，这两个类型的量与其他量的类型不一致时，会尝试把所有的量转换成一致的类型，成功就不报错，不成功就报错。
	//所以a和b可以运算而不报错，但是d中string类型的"3"和整型的数字无法转换成统一的类型，所以出现了类型不匹配的错误。
}