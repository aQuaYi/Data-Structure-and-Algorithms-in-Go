//本程序没有严谨地考虑，0和别的数的互质问题。按照书本上gcd的定义，简单认为0只与1互质。
//本程序复制了1.1.11程序的代码。
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := makeBool2d(20)
	b.print()
}

func gcd(a, b int) (result int) {
	if a < b {
		a, b = b, a
	}

	if b == 0 {
		result = a
		return
	}

	return gcd(b, a%b)
}

type bool2d [][]bool

func makeBool2d(n int) bool2d {
	b := bool2d{}
	for i := 0; i < n; i++ {
		b = append(b, []bool{})
		for j := 0; j < n; j++ {
			if gcd(i, j) == 1 {
				b[i] = append(b[i], true)
			} else {
				b[i] = append(b[i], false)
			}
		}
	}
	return b
}

func (b bool2d) print() {
	rowNumWid := width(len(b))
	rowNumFormat := "%" + strconv.Itoa(rowNumWid) + "d "

	//打印左上角的空白地带
	printSpace(rowNumWid + 1)

	//打印列号
	for i := 0; i < len(b[0]); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	for i := 0; i < len(b); i++ {
		fmt.Printf(rowNumFormat, i)      //打印行号
		for j := 0; j < len(b[i]); j++ { //打印矩阵内容
			printBool(b[i][j])
			printSpace(width(j))
		}
		fmt.Println()
	}
}

func width(n int) int {
	i := 0
	for n > 0 {
		n /= 10
		i++
	}
	return i
}

func printBool(b bool) {
	s := " "
	if b {
		s = "*"
	}
	fmt.Print(s)
}

func printSpace(width int) {
	if width == 0 { //不添加这一个话，0列和1列的星号会挨在一起。
		width++
	}
	for i := 0; i < width; i++ {
		fmt.Print(" ")
	}
}
