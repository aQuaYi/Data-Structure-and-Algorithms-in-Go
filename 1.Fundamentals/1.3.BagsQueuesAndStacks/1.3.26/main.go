package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("程序开始")
	rand.Seed(time.Now().UnixNano())
	randKey := 4
	n := 30
	fmt.Println("n=", n)
	key := rand.Intn(randKey)
	fmt.Println("key=", key)

	first := newNode(0)
	last := first

	for i := 1; i < n; i++ {
		newND := newNode(rand.Intn(randKey))
		last.next = newND
		last = newND
	}

	temp := first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()

	first = remove(first, key)

	temp = first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()
}

func remove(first *node, key interface{}) *node {
	for first.item == key { //处理链表一开头就出现key的情况。
		first = first.next
	}

	move := first

	for move.next != nil { //处理链表中间出现key的情况。
		if move.next.item == key {
			move.next = move.next.next
			continue
		}
		move = move.next
	}

	return first
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
