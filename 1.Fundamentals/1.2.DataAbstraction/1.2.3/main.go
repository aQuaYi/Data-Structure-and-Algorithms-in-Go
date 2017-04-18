package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("请输入3个参数，正整数N和属于[0,1]的float64数min和max，要求min<max。")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("无法获取正整数N", err)
		os.Exit(1)
	}
	min, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("无法获取下限min", err)
		os.Exit(1)
	}
	max, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil || min >= max {
		fmt.Println("无法获取上限max，或者min>=max。", err)
		os.Exit(1)
	}

}
