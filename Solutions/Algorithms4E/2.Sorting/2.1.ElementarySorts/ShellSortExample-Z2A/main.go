package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("程序开始")

	str := "ZYXWVUTSRQPONMLKJIHGFEDCBA"
	ss := strings.Split(str, "")
	n := len(ss)
	h := 1
	compareCount := 0
	exchangeCount := 0
	hFactor := 3
	for h < n/hFactor {
		h = hFactor*h + 1
	}

	printS(ss, -1, -1, "", "")
	fmt.Println("--------------------------------------------")
	for h >= 1 {
		fmt.Println("h=", h, "------------------------")
		for i := h; i < n; i++ {
			for j := i; j >= h; j -= h {
				printS(ss, j, j-h, "_", "_") //两个相互比较
				if ss[j] < ss[j-h] {
					exchangeCount++
					ss[j], ss[j-h] = ss[j-h], ss[j]
					printS(ss, j, j-h, "[", "]") //两个已经交换位置的元素
					fmt.Println()
					continue
				}
				compareCount++
				break
			}
		}
		h = h / hFactor
	}
	fmt.Println("--------------------------------------------")
	printS(ss, -1, -1, "", "")
	fmt.Println("总比较次数", compareCount)
	fmt.Println("总的交换次数", exchangeCount)
}
func printS(ss []string, a, b int, sl, sr string) {
	for i, v := range ss {
		if i == a || i == b {
			fmt.Printf("%s%s%s", sl, v, sr)
			continue
		}
		fmt.Printf(" %s ", v)
	}
	fmt.Println()
}
