package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(count(histogram(randIS(10000, 10), 10)))
	fmt.Println(count(histogram(randIS(10000, 12), 10)))
}

func count(is []int) int {
	result := 0
	for _, v := range is {
		result += v
	}
	return result
}

func histogram(is []int, m int) []int {
	result := make([]int, m)
	for _, v := range is {
		if v < m {
			result[v]++
		}
	}
	return result
}

func randIS(l, top int) []int {
	result := []int{}
	for i := 0; i < l; i++ {
		result = append(result, rand.Intn(top))
	}
	return result
}
