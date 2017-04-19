package main

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	a := NewAccumulator()
	rand.Seed(time.Now().UnixNano())
	xs := []float64{}
	bxs := []*big.Float{}
	bsum := big.NewFloat(0)
	sum := 0.
	n := 1000.
	for i := 0.; i < n; i++ {
		x := rand.Float64() * 100
		a.AddDataValue(x)
		xs = append(xs, x)
		sum += x
		bxs = append(bxs, big.NewFloat(x))
		bsum.Add(x *big.Float, y *big.Float)
		bsum.Add(x *big.Float, y *big.Float)
	}

	big.NewFloat()
}

//Accumulator 是累加器
type Accumulator struct {
	m, s, n float64
}

//NewAccumulator 创造了一个*Accumulator实例
func NewAccumulator() *Accumulator {
	return &Accumulator{}
}

//AddDataValue 向累加器中添加数据
func (a *Accumulator) AddDataValue(x float64) {
	a.n++
	a.s = a.s + (a.n-1)/a.n*(x-a.m)*(x-a.m)
	a.m = a.m + (x-a.m)/a.n
}

//Mean 返回累加器的均值
func (a *Accumulator) Mean() float64 {
	return a.m
}

//Var 返回累加器的方差
func (a *Accumulator) Var() float64 {
	return a.s / (a.n - 1)
}

//StdDev 返回累加器的标准差
func (a *Accumulator) StdDev() float64 {
	return math.Sqrt(a.Var())
}
