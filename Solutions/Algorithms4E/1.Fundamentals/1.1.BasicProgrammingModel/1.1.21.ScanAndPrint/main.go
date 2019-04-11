package main

import (
	"fmt"
)

func main() {
	fmt.Println("输入格式：\nTim 1 2")
	fmt.Println("输入end，结束输入。")
	records := []record{}
	n, f, s := "", 0, 0
	for {
		fmt.Scanf("%s %d %d", &n, &f, &s)
		if n == "end" {
			break
		}
		if s == 0 {
			fmt.Println("第二个数不能为0，请重新输入。")
			continue
		}
		if n == "" {
			continue
		}
		records = append(records, record{
			name:   n,
			first:  f,
			second: s,
			result: float64(f) / float64(s),
		})
		n, f, s = "", 0, 1
	}
	for _, r := range records {
		fmt.Println(r)
	}
}

type record struct {
	name          string
	first, second int
	result        float64
}

func (r record) String() string {
	return fmt.Sprintf("%s\t %d %d %.3f", r.name, r.first, r.second, r.result)
}
