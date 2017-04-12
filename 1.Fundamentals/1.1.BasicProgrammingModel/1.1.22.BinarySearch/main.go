package main

import "fmt"

func main() {
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
	}

	fmt.Println(-1, rank(-1, a))
	fmt.Println(78, rank(78, a))
	fmt.Println(1024, rank(1024, a))
}

func rank(key int, a []int) int {
	return rankRecur(key, a, 0, len(a)-1, 0)
}

func rankRecur(key int, a []int, lo, hi, depth int) int {
	spacePrint(depth)
	fmt.Println(lo, hi)
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	switch {
	case a[mid] > key:
		return rankRecur(key, a, lo, mid-1, depth+1)
	case a[mid] < key:
		return rankRecur(key, a, mid+1, hi, depth+1)
	default:
		return mid
	}
}

func spacePrint(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
}
