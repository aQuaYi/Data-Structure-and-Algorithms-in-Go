package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/fogleman/gg"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("请在程序名称后，添加一个整数N和两个浮点数l，r")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("无法获取正整数N:", err)
		os.Exit(1)
	}
	l, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("无法获取正确的l,l应该是一个浮点数。", err)
		os.Exit(1)
	}

	r, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil || l >= r {
		fmt.Println("无法获取正确的r,r应该是一个浮点数，且大于l。", err)
		os.Exit(1)
	}

	count := make([]int, n)
	numbers := make([]float64, n)
	for i := 0; i < n; i++ {
		numbers[i] = l + (r-l)*float64(i+1)/float64(n)
	}

	str := ""
	for {
		if _, err := fmt.Scanln(&str); err != nil {
			if err == io.EOF {
				fmt.Println("读取结束。")
				break
			}
			fmt.Println(str, "无法读取", err)
		}
		t, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("读取的内容无法转换成float64", str, err)
			continue
		}
		if t <= l || r <= t {
			continue
		}

		i := rank(t, numbers)
		count[i]++
	}

	const (
		w      = 80
		wBlank = 40
		h      = 1000
		hBlank = 100
	)
	W := float64(w*n + wBlank*(n+1))
	H := float64(h + hBlank*2)

	dc := gg.NewContext(int(W), int(H))
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	y := float64(h + hBlank)
	max := float64(maxOf(count))
	if err := dc.LoadFontFace("/Library/Fonts/scp.ttf", 36); err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		if count[i] == 0 {
			continue
		}
		x := wBlank*float64(i+1) + w*float64(i)
		hi := -1.0 * h * float64(count[i]) / max
		dc.DrawRectangle(x, y, w, hi)
		dc.Fill()
		dc.DrawStringAnchored(strconv.Itoa(count[i]), x, y, 0, 1.5)
	}
	dc.SavePNG("out.png")
}

func maxOf(a []int) int {
	result := 0
	for _, v := range a {
		if v > result {
			result = v
		}
	}
	return result
}

// 以下是借用1.1.29的代码
func rank(key float64, a []float64) int {
	return rankRecur(key, a, 0, len(a)-1)
}

func rankRecur(key float64, a []float64, lo, hi int) int {
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
