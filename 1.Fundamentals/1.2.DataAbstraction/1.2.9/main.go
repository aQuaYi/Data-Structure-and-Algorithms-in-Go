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

	"Algorithms-in-Golang/counter"
)

func main() {
	bs := counter.New("BinarySearch")

	whiteList, err := readInts(os.Args[1])
	if err != nil {
		fmt.Println("读取白名单失败：", err)
		os.Exit(1)
	}
	opt := os.Args[2]
	if !(opt == "+" || opt == "-") {
		fmt.Println("第二个参数应为“+”或者“-”")
		os.Exit(1)
	}
	if opt == "+" {
		fmt.Println("接下来，会打印出*不*在白名单上的值。")
	} else {
		fmt.Println("接下来，会打印出在白名单上的值。")
	}

	sort.Ints(whiteList)

	str := ""
	i := 0
	for {
		if _, err := fmt.Scanln(&str); err != nil {
			if err == io.EOF {
				fmt.Println("读取结束。")
				break
			}
			fmt.Println(str, "无法读取", err)
		}
		t, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("读取的内容无法转换成整数", str, err)
		}
		r := rank(t, whiteList, bs)

		switch {
		case opt == "+" && r == -1:
			fmt.Println(t)
		case opt == "-" && r != -1:
			i++
			fmt.Println(t)
		}
		str = ""
	}
	fmt.Printf("程序一共输出了%d个不在WhiteList中的数字\n", i)
	fmt.Printf("一共在WhiteList中找到了%d个数字\n", bs.Tally())
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

func rank(key int, a []int, c *counter.Counter) int {
	return rankRecur(key, a, 0, len(a)-1, c)
}

func rankRecur(key int, a []int, lo, hi int, c *counter.Counter) int {
	if lo > hi {
		return -1
	}
	mid := (lo + hi) / 2
	switch {
	case a[mid] > key:
		return rankRecur(key, a, lo, mid-1, c)
	case a[mid] < key:
		return rankRecur(key, a, mid+1, hi, c)
	default:
		c.Increment()
		return mid

	}
}
