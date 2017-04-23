//我对题目中说，“只能使用一个Node类型的实例变量（last）”的理解是：queue的属性只能有last。
package main

import (
	"fmt"

	"./circleQueue"
)

func main() {
	fmt.Println("程序开始")
	q := circleQueue.New()
	n := 12
	fmt.Printf("将会输出[0,%d]中的所有整数\n", n-1)
	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}
	c := q.Iterator()
	for c.HasNext() {
		fmt.Printf("%v ", c.Next())
	}
}
