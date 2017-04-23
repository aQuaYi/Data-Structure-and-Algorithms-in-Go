package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("程序开始")
	rand.Seed(time.Now().UnixNano())
	n := 10
	fmt.Println("n=", n)

	first := newNode(2)
	last := first

	for i := 1; i < n; i++ {
		newND := newNode(rand.Intn(n * 3))
		last.next = newND
		last = newND
	}

	temp := first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()

	maxInt := max(first)

	fmt.Println("max=", maxInt)
	fmt.Println("空链条的情况")
	fmt.Println("max(nil)=", max(nil))
}

func max(first *node) int {
	if first == nil {
		return 0
	}
	fint := first.item.(int)

	return bigger(fint, max(first.next))
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type node struct {
	item interface{}
	next *node
}

func newNode(item interface{}) *node {
	return &node{
		item: item,
	}
}
