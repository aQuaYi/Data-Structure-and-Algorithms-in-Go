package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("程序开始")
	rand.Seed(time.Now().UnixNano())

	n := 20
	fmt.Println("n=", n)
	k := rand.Intn(n * 2)
	fmt.Println("k=", k)

	insertNode := newNode(99)

	first := newNode(0)
	last := first
	var x *node

	for i := 1; i < n; i++ {
		newND := newNode(i)
		last.next = newND
		last = newND
		if i == k {
			x = newND
		}
	}

	temp := first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()

	if insertNode != nil && x != nil {
		insertAfter(x, insertNode)
	}

	temp = first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()
}

func insertAfter(x, insert *node) {
	insert.next = x.next
	x.next = insert
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
