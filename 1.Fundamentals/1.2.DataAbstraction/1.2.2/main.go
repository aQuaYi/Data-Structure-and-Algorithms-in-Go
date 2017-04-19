package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type line struct {
	begin, end float64
}

func (l line) String() string {
	return fmt.Sprintf("[%.2f,%.2f]", l.begin, l.end)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("请输入一个正整数参数N")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("无法获取正整数参数N:", err)
		os.Exit(1)
	}
	fmt.Println("请输入float64数对，使用“,”分隔")
	ls := []line{}
	for n > 0 {
		str := ""
		fmt.Printf("还剩%d个\n", n)
		if _, err := fmt.Scanln(&str); err != nil {
			fmt.Println(str, "无法被读取", err)
			continue
		}
		ss := strings.Split(str, ",")
		b, err := strconv.ParseFloat(ss[0], 64)
		if err != nil {
			fmt.Println(ss[0], "无法转换成float64:", err)
			fmt.Println("输入float64数对，请使用“,”分隔")
			continue
		}
		e, err := strconv.ParseFloat(ss[1], 64)
		if err != nil {
			fmt.Println(ss[1], "无法转换成float64:", err)
			fmt.Println("输入float64数对，请使用“,”分隔")
			continue
		}
		if b > e {
			b, e = e, b
		}
		ls = append(ls, line{
			begin: b,
			end:   e,
		})

		n--
	}
	fmt.Println("输入结束。")
	fmt.Println("以下间隔相交了：")
	for i := 0; i < len(ls); i++ {
		for j := i + 1; j < len(ls); j++ {
			if intersect(ls[i], ls[j]) {
				fmt.Println(ls[i], ls[j])
			}
		}
	}

}

func intersect(a, b line) bool {
	min := math.Min(a.begin, b.begin)
	max := math.Max(a.end, b.end)
	if (max - min) >= (a.end - a.begin + b.end - b.begin) {
		return false
	}
	return true
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
