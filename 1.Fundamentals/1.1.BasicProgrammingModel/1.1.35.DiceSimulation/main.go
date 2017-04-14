package main

import (
	"fmt"
	"math"
	"math/rand"
)

const SIDES = 6

type doubleDice [2*SIDES + 1]float64

func main() {
	dist := doubleDice{}
	for i := 1; i <= SIDES; i++ {
		for j := 1; j <= SIDES; j++ {
			dist[i+j] += 1.
		}
	}
	for i := range dist {
		dist[i] /= 36.
		fmt.Println(i, dist[i])
	}

	testDist := doubleDice{}
	tempDist := doubleDice{}
	n := 0
	for !finish(dist, testDist) {
		n++
		i := rand.Intn(6) + 1
		j := rand.Intn(6) + 1
		tempDist[i+j]++
		testDist = probabilityOf(tempDist)
	}

	fmt.Println(n)
	for i := range testDist {
		fmt.Printf("%d %.4f %.4f\n", i, dist[i], testDist[i])
	}
}

func finish(a, b doubleDice) bool {
	for i, v := range a {
		if math.Abs(b[i]-v) > 0.001 {
			return false
		}
	}
	return true
}

func probabilityOf(temp doubleDice) doubleDice {
	result := doubleDice{}
	sum := 0.
	for _, v := range temp {
		sum += v
	}
	for i, v := range temp {
		result[i] = v / sum
	}
	return result
}
