/*
这一题最大的难度在于，当行号和列号的宽度慢慢变宽时，保持星号的对齐。
*/
package main

import "fmt"
import "strconv"
import "math/rand"

func main() {
	b2d := randBool2d(40, 20)
	b2d.print()
}

type bool2d [][]bool

func randBool2d(m, n int) bool2d {
	b := bool2d{}
	for i := 0; i < m; i++ {
		b = append(b, []bool{})
		for j := 0; j < n; j++ {
			if rand.Float64() > 0.5 {
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
	if width == 0 {
		width++
	}
	for i := 0; i < width; i++ {
		fmt.Print(" ")
	}
}
