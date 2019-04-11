package main

import "fmt"

func main() {
	fmt.Println("reverse长度为0的链")
	fmt.Println(reverse(nil))

	fmt.Println("reverse长度为１的链")
	first := newNode(0)
	fmt.Println(first.item)
	fmt.Println(reverse(first).item)
	n := 20
	fmt.Printf("reverse长度为%d的链\n", n)

	last := first
	for i := 1; i < n; i++ {
		newLast := newNode(i)
		last.next = newLast
		last = newLast
	}

	temp := first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()

	first = reverse(first)
	temp = first
	for temp != nil {
		fmt.Printf("%v ", temp.item)
		temp = temp.next
	}
	fmt.Println()
}
func reverse(x *node) *node {
	first := x
	var reverse *node

	for first != nil {
		second := first.next
		first.next = reverse
		reverse = first
		first = second
	}

	return reverse
}

// func reverse(first *node) *node {
// 	if first == nil {
// 		return nil
// 	}
// 	if first.next == nil {
// 		return first
// 	}

// 	f, fn, fnn := first, first.next, first.next.next
// 	first.next = nil
// 	for fnn.next != nil {
// 		fn.next = f
// 		f, fn, fnn = fn, fnn, fnn.next

// 	}

// 	fn.next = f
// 	fnn.next = fn

// 	return fnn
// }

type node struct {
	item interface{}
	next *node
}

func newNode(item interface{}) *node {
	return &node{
		item: item,
	}
}

func (n *node) nextTo(anther *node) *node {
	n.next = anther
	return n
}
