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
	// maxInt := max(nil)

	fmt.Println("max=", maxInt)
}

func max(first *node) int {
	if first == nil {
		return 0
	}

	max := first.item.(int)

	for first.next != nil { //处理链表中间出现key的情况。
		temp := first.next.item.(int)
		if temp > max {
			max = temp
		}
		first = first.next
	}

	return max
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
