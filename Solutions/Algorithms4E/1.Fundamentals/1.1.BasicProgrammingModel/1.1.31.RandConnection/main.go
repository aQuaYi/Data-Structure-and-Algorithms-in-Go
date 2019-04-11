package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/fogleman/gg"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("请在程序名称后，添加一个整数N和一个属于[0,1)的小数p")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("无法获取正整数N:", err)
		os.Exit(1)
	}
	p, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil || p < 0 || 1 <= p {
		fmt.Println("无法获取正确的小数。", err)
		os.Exit(1)
	}

	const S = 1000
	const X = S / 2
	const Y = S / 2
	const R = S * 2 / 5.0
	const r = 10
	degree := 360.0 / float64(n)
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = X + R*math.Cos(gg.Radians(degree*float64(i)))
		y[i] = Y + R*math.Sin(gg.Radians(degree*float64(i)))
	}

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	for i := 0; i < n; i++ {
		dc.DrawCircle(x[i], y[i], r)
	}
	dc.Fill()

	dc.SetLineWidth(3)
	dc.SetRGBA(0, 0, 0, 0.3)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if rand.Float64() > p {
				dc.DrawLine(x[i], y[i], x[j], y[j])
			}
		}
	}
	dc.Stroke()
	dc.SavePNG("out.png")
}
