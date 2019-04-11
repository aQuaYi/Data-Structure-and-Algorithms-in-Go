package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("请输入参数，数组的长度M和打乱的次数N")
		os.Exit(1)
	}
	m, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("无法获取输入的数组长度Ｍ:", err.Error())
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("无法获取打乱的次数Ｎ:", err.Error())
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, m)
	}

	for k := 0; k < n; k++ {
		a := newInts(m)
		a = shuffle(a)
		for i := 0; i < m; i++ {
			result[i][a[i]]++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%7d ", result[i][j])
		}
		fmt.Println()
	}
}

func newInts(m int) []int {
	result := make([]int, m)
	for i := 0; i < m; i++ {
		result[i] = i
	}
	return result
}

func shuffle(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		r := rand.Intn(n) //和36题相比，只有这里不同。
		a[i], a[r] = a[r], a[i]
	}
	return a
}
