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

func reverse(first *node) *node {
	if first == nil {
		return nil //仅当first为nil时，会执行这一条语句。
	}
	if first.next == nil {
		//当first不为nil时，这是最简单的情况。
		//这里返回的first已经是链条的末尾了。
		return first
	}

	second := first.next
	rest := reverse(second)
	second.next = first
	first.next = nil
	return rest

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

func (n *node) nextTo(anther *node) *node {
	n.next = anther
	return n
}
