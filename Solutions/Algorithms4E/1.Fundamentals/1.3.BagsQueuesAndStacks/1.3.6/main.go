package main

import (
	"Algorithms-in-Golang/queue"
	"Algorithms-in-Golang/stack"
	"fmt"
)

func main() {
	fmt.Println("程序开始")
	q := queue.New()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	c := q.Iterator()
	for c.HasNext() {
		fmt.Printf("%v ", c.Next())
	}

	fmt.Println()

	s := stack.New()
	for !q.IsEmpty() {
		s.Push(q.Dequeue())
	}
	for !s.IsEmpty() {
		q.Enqueue(s.Pop())
	}

	c = q.Iterator()
	for c.HasNext() {
		fmt.Printf("%v ", c.Next())
	}
}
