package main

import (
	"fmt"
)

func main() {
	// 因为以下计算过程是int类型的，所有每一步都会舍弃小数点部分。
	fmt.Println((0 + 15) / 2)
	// 这一题应该是考查极大的数的末尾精度问题，但是Golang没有出错，得出了正确答案。
	fmt.Println(2.0e-6 * 100000000.1)
	// 在Golang中 &&　与　||　是同级的。所以，遵循从左往右的循序。
	fmt.Println(true && false || true && true)
	fmt.Println(false && true || true) // 答案是true
}
