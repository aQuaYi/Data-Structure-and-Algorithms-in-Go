package main

import (
	"fmt"
)

func main() {
	// 因为以下计算过程是int类型的，所有每一步都会舍弃小数点部分。
	fmt.Println((0 + 15) / 2)
	fmt.Println(2.0e-6 * 100000000.1)
	fmt.Println(true && false || true && true)
}
