package main

import (
	"fmt"
	"time"
)

var b map[string]float64

func main() {
	b = make(map[string]float64)
	beginTime := time.Now()
	fmt.Println(binomal(100, 50, 0.5))
	fmt.Println("耗时", time.Since(beginTime))
}

func binomal(n, k int, p float64) float64 {
	if n == 0 && k == 0 {
		return 1
	}

	if n < 0 || k < 0 || n < k {
		return 0
	}

	key := fmt.Sprintf("%d-%d-%f", n, k, p)
	if v, ok := b[key]; ok {
		return v
	}

	v := (1-p)*binomal(n-1, k, p) + p*binomal(n-1, k-1, p)
	b[key] = v
	return v

}
