package main

import "fmt"

func main() {
	fmt.Println("程序开始")
	a := []int{4, 3, 2, 1, 0, 9, 8, 7, 6, 5}
	b := []int{4, 6, 8, 7, 5, 3, 2, 9, 0, 1}
	c := []int{2, 5, 6, 7, 4, 8, 9, 2, 1, 0}
	d := []int{4, 3, 2, 1, 0, 5, 6, 7, 8, 9}
	e := []int{1, 2, 3, 4, 5, 6, 9, 8, 7, 0}
	f := []int{0, 4, 6, 5, 3, 8, 1, 7, 2, 9}
	g := []int{1, 4, 7, 9, 8, 6, 5, 3, 0, 2}
	h := []int{2, 1, 4, 3, 6, 5, 8, 7, 9, 0}
	checkStackPop(a)
	checkStackPop(b)
	checkStackPop(c)
	checkStackPop(d)
	checkStackPop(e)
	checkStackPop(f)
	checkStackPop(g)
	checkStackPop(h)
}

func checkStackPop(a []int) {
	for i := len(a) - 1; i > 1; i-- {
		if a[i-1] > a[i] {
			continue
		}
		for j := 0; j < i; j++ {
			if a[j] > a[i] {
				fmt.Println("此出栈队列不可能：", a)
				fmt.Printf("    因为%d在%d前面了\n", a[j], a[i])
				return
			}
		}
	}
	fmt.Println("此出栈队列可能：", a)
}
