package main

import (
	"fmt"
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
		x := rand.Float64()

		a.AddDataValue(x)

		xs = append(xs, x)
		sum += x

		bigx := big.NewFloat(x)
		bxs = append(bxs, bigx)
		bsum.Add(bsum, bigx)
	}
	mean := sum / n
	bmean := bsum.Mul(bsum, big.NewFloat(1/(n-1)))

	vsum := 0.
	for _, v := range xs {
		vsum += (v - mean) * (v - mean)
	}
	varv := vsum / (n - 1)
	stddev := math.Sqrt(varv)

	bvsum := big.NewFloat(0)
	for _, v := range bxs {
		v.Sub(v, bmean)
		bvsum.Add(bvsum, v.Mul(v, v))
	}
	bvar := bvsum.Mul(bvsum, big.NewFloat(1/(n-1)))
	bvar2f, _ := bvar.Float64() //math/big 没有提供sqrt()，只能先转换成float64再开方。
	bstddev := math.Sqrt(bvar2f)

	fmt.Println("a.Var() =", a.Var())
	fmt.Println("    var =", varv)
	fmt.Println("   bvar =", bvar)
	fmt.Println("a.StdDev() =", a.StdDev())
	fmt.Println("    stddev =", stddev)
	fmt.Println("    bstddev=", bstddev)
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
	a.s += (a.n - 1) / a.n * (x - a.m) * (x - a.m)
	a.m += (x - a.m) / a.n
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
