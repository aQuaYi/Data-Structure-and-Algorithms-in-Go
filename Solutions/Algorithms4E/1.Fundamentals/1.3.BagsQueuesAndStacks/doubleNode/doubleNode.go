package doubleNode

import (
	"sync"
)

//Chain 是基于doubleNode的双向链接口
type Chain interface {
	AddFirst(interface{})
	AddLast(interface{})
	RemoveFirst() interface{}
	RemoveLast() interface{}

	//InsertBefore 是在结点的first边插入
	InsertBefore(interface{}, interface{})

	//InsertAfter 是在结点的last边插入
	InsertAfter(interface{}, interface{})

	Remove(interface{})

	Size() int
}

type chain struct {
	first, last *node
	n           int
	sync.RWMutex
}

//New 返回一个空的chain,作为Chain接口
func New() Chain {
	return &chain{}
}

func (c *chain) Size() int {
	c.RLock()
	defer c.RUnlock()
	return c.n
}

func (c *chain) AddFirst(a interface{}) {
	c.Lock()
	defer c.Unlock()
	newFirst := newNode(a)
	if c.n == 0 {
		c.first = newFirst
		c.last = newFirst
		c.n++
		return
	}

	c.first.first = newFirst
	newFirst.last = c.first
	c.first = newFirst
	c.n++
}

func (c *chain) AddLast(a interface{}) {
	c.Lock()
	defer c.Unlock()
	newLast := newNode(a)
	if c.n == 0 {
		c.first = newLast
		c.last = newLast
		c.n++
		return
	}

	c.last.last = newLast
	newLast.first = c.last
	c.last = newLast
	c.n++
}

func (c *chain) RemoveFirst() interface{} {
	c.Lock()
	defer c.Unlock()

	if c.n == 0 {
		return nil
	}

	result := c.first

	if c.n == 1 {
		c.first, c.last = nil, nil
	} else {
		result.last.first = nil
		c.first = result.last
	}

	c.n--
	return result.item
}

func (c *chain) RemoveLast() interface{} {
	c.Lock()
	defer c.Unlock()

	if c.n == 0 {
		return nil
	}

	result := c.last

	if c.n == 1 {
		c.first, c.last = nil, nil
	} else {
		result.first.last = nil
		c.last = result.first
	}

	c.n--
	return result.item
}

func (c *chain) InsertBefore(a interface{}, key interface{}) {
	if n == c.first {
		c.AddFirst(a)
	}

	c.Lock()
	defer c.Unlock()

	newNode := newNode(a)
	f := n.first
	f.last = newNode
	newNode.last = n
	n.first = newNode
	newNode.first = f
}

func (c *chain) InsertAfter(a interface{}, key interface{}) {
	if n == c.last {
		c.AddLast(a)
	}

	c.Lock()
	defer c.Unlock()

	newNode := newNode(a)
	l := n.last
	n.last = newNode
	newNode.last = l
	l.first = newNode
	newNode.first = n
}

func (c *chain) Remove(key interface{}) {

	if c.n == 0 {
		return
	}

	if c.first.Equal(key) {
		c.RemoveFirst()
		return
	}

	if c.last.Equal(key) {
		c.RemoveLast()
		return
	}

	c.Lock()
	defer c.Unlock()

	if c.first == n {
		c.first = c.first.last
		c.first.first = nil
		return
	}

	if c.last == n {
		c.last = c.last.first
		c.last.last = nil
		return
	}

	f, l := n.first, n.last
	f.last = l
	l.first = f
	return
}
