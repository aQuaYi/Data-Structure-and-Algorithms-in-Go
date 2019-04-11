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

	first := newNode(0)
	last := first
	var removeNode *node

	for i := 1; i < n; i++ {
		newND := newNode(i)
		last.next = newND
		last = newND
		if i == k {
			removeNode = newND
		}
	}

	temp := first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()

	if removeNode != nil {
		removeAfter(removeNode)
	}

	temp = first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()
}

func removeAfter(rn *node) {
	if rn.next != nil {
		rn.next = rn.next.next
	}
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
