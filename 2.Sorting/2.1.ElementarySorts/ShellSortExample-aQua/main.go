package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("程序开始")

	max := 1000000

	ss := make([]int, max)

	for i := 0; i < max; i++ {
		ss[i] = rand.Intn(max)
	}
	fmt.Println("数据准备完毕")
	n := len(ss)
	h := 1
	count := 0
	eConunt := 0
	hn := 3
	for h < n/hn {
		h = hn*h + 1
	}
	beginTime := time.Now()
	for h >= 1 {
		fmt.Println("h=", h, time.Since(beginTime))
		for i := h; i < n; i++ {
			for j := i; j >= h; j -= h {
				if ss[j] < ss[j-h] {
					ss[j], ss[j-h] = ss[j-h], ss[j]
					eConunt++
				}
				count++
			}
		}
		h = h / hn
	}

	fmt.Println("总比较次数", count)
	fmt.Println("总交换次数", eConunt)
}
