package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	whiteList, err := readInts(os.Args[1])
	if err != nil {
		fmt.Println("读取白名单失败：", err)
		os.Exit(1)
	}

	sort.Ints(whiteList)

	for i := 0; i < len(whiteList); i++ {
		fmt.Println(i+1, whiteList[i])
	}
	fmt.Println("rank(9)=", rank(9, whiteList))
	fmt.Println("rank(14)=", rank(14, whiteList))
	fmt.Println("rank(20)=", rank(20, whiteList))
	fmt.Println("rank(50)=", rank(50, whiteList))
	fmt.Println("rank(51)=", rank(51, whiteList))
	fmt.Println("rank(98)=", rank(98, whiteList))
	fmt.Println("rank(101)=", rank(101, whiteList))

	fmt.Println("count(12)=", count(12, whiteList))
	fmt.Println("count(20)=", count(20, whiteList))
	fmt.Println("count(50)=", count(50, whiteList))
	fmt.Println("count(98)=", count(98, whiteList))
	fmt.Println("count(99)=", count(99, whiteList))
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
			if number, err = strconv.Atoi(string(line)); err != nil {
				fmt.Println("无法转换成数字")
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
		return lo
	}
	mid := (lo + hi) / 2
	switch {
	case a[mid] >= key:
		return rankRecur(key, a, lo, mid-1)
	default:
		return rankRecur(key, a, mid+1, hi)
	}
}

func count(key int, a []int) int {
	result := 0
	begin := rank(key, a)
	for i := begin; i < len(a); i++ {
		if a[i] == key {
			result++
		} else {
			break
		}
	}
	return result
}
