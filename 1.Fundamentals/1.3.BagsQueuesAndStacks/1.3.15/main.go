//我输出的是int,原理是一样的。
package main

import (
	"Algorithms-in-Golang/queue"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("程序开始")

	k := 10
	q := queue.New()
	r := rand.Intn(1000)
	for i := 0; i < k+r; i++ {
		q.Enqueue(i)
	}
	fmt.Printf("%d=", r)

	n := q.Size()
	c := q.Iterator()
	for c.HasNext() {
		result := c.Next().(int)
		if n == k {
			fmt.Printf("%v\n", result)
		}
		n--
	}
}
