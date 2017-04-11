package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i2d := randInt2d(5, 14)
	i2d.print()
	i2d.trans().print()
}

type int2d [][]int

func randInt2d(m, n int) int2d {
	i2d := int2d{}
	for i := 0; i < m; i++ {
		i2d = append(i2d, []int{})
		for j := 0; j < n; j++ {
			i2d[i] = append(i2d[i], rand.Intn(1000))
		}
	}
	return i2d
}

func (i2d int2d) print() {
	format := "%4d "
	for i := range i2d {
		for j := range i2d[i] {
			fmt.Printf(format, i2d[i][j])
		}
		fmt.Println()
	}
}

func (i2d int2d) trans() int2d {
	newI2D := int2d{}
	for j := range i2d[0] {
		newI2D = append(newI2D, []int{})
		for i := range i2d {
			newI2D[j] = append(newI2D[j], i2d[i][j])
		}
	}
	return newI2D
}
