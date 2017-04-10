package main

import (
	"fmt"
)

func main() {
	var a []int
	for i := 0; i < 10; i++ {
		a[i] = i * i
		//这个程序的Golang版本也会出错，但是出错原因不一样。
		//var a []int 后，给a分配的是一个长度为0的底层数组。
		//所以，当a[0]=0赋值的时候，会报错“panic: runtime error: index out of range”，超出了a的长度范围。
		//上句改为下句，就可以执行了
		//a = append(a, i*i)
	}
	fmt.Println(a)
}
