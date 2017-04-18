package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/fogleman/gg"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("请在程序名称后，添加一个整数N")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("无法获取正整数N:", err)
		os.Exit(1)
	}

	const S = 1000
rand.Seed(time.Now().UnixNano())
	ps := []gg.Point{}
	for i := 0; i < n; i++ {
		ps = append(ps, gg.Point{
			X: rand.Float64() * S,
			Y: rand.Float64() * S,
		})
	}

	a, b := closest(ps)

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	for i := 0; i < n; i++ {
		dc.DrawPoint(ps[i].X, ps[i].Y, 3)
	}
	dc.Fill()

	dc.SetLineWidth(3)
	dc.SetRGBA(0, 0, 0, 0.3)
	dc.DrawLine(a.X, a.Y, b.X, b.Y)
	dc.Stroke()
	dc.SavePNG("out.png")
}

func closest(ps []gg.Point) (a, b gg.Point) {
	d := math.MaxFloat64
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			dij := distance(ps[i], ps[j])
			if dij < d {
				d = dij
				a, b = ps[i], ps[j]
			}
		}
	}
	return
}

func distance(a, b gg.Point) float64 {
	x := a.X - b.X
	y := a.Y - b.Y
	return math.Sqrt(x*x + y*y)
}
