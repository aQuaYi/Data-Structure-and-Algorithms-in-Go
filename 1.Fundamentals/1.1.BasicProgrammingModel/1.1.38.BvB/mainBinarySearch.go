package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	whiteList, err := readInts(os.Args[1])
	if err != nil {
		fmt.Println("读取白名单失败：", err)
		os.Exit(1)
	}

	sort.Ints(whiteList)
	beginTime := time.Now()
	count := 0

	for {
		str := ""
		if _, err := fmt.Scanln(&str); err != nil {
			if err == io.EOF {
				fmt.Println("读取结束。")
				break
			}
			fmt.Println(str, "无法读取", err)
			continue
		}

		t, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("读取的内容无法转换成整数", str, err)
		}
		r := rank(t, whiteList)
		if r != -1 {
			count++
		}
	}
	fmt.Printf("一共找到%d个数字在白名单内\n", count)
	fmt.Println("总耗时:", time.Since(beginTime))
}

func readInts(filename string) ([]int, error) {
	//打开文件，并进行相关处理
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("打开文件失败:" + err.Error())
	}
	//文件关闭
	defer f.Close()

	result := []int{}
	number := 0
	//将文件作为一个io.Reader对象进行buffered I/O操作
	br := bufio.NewReader(f)
	for {
		//每次读取一行
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else {
			str := strings.Replace(string(line), " ", "", -1)
			if number, err = strconv.Atoi(str); err != nil {
				fmt.Println(line, str, "无法转换成数字")
				continue
			}
			result = append(result, number)
		}
	}
	return result, nil
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
