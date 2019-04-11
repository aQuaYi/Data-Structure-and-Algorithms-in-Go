package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("请添加实验重复次数T")
		os.Exit(1)
	}

	t, err := strconv.Atoi(os.Args[1])
	if err != nil || t <= 0 {
		fmt.Println("无法把T转换为正整数：", err)
	}

	ns := []int{1e3, 1e4, 1e5, 1e6}

	for _, n := range ns {
		count := 0.
		for i := 0; i < t; i++ {
			sA := makeInts(n)
			sort.Ints(sA)
			sB := makeInts(n)
			for _, v := range sB {
				if rank(v, sA) != -1 {
					count++
				}
			}
		}
		fmt.Printf("======%d======\n", n)
		fmt.Println(count / float64(t))
	}
}

func makeInts(n int) []int {
	result := make([]int, n)
	for i := range result {
		result[i] = rand.Intn(1000000)
	}
	return result
}

func rank(key int, a []int) int {
	return rankRecur(key, a, 0, len(a)-1)
}

func rankRecur(key int, a []int, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	switch {
	case a[mid] > key:
		return rankRecur(key, a, lo, mid-1)
	case a[mid] < key:
		return rankRecur(key, a, mid+1, hi)
	default:
		return mid
	}
}
