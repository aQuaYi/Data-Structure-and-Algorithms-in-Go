package bag

//Iterator 是为了实现Bag的迭代功能而实现的接口。
type Iterator interface {
	//Item 返回迭代器当前的值
	Item() interface{}

	//Next
	Next() interface{}

	// HasNext checks whether there is a new value available.
	HasNext() bool
}

type chain struct {
	node *node
}

func (c *chain) Item() interface{} {
	return c.node.item
}

func (c *chain) Next() interface{} {
	c.node = c.node.next
	return c.node.item
}

func (c *chain) HasNext() bool {
	if c.node.next != nil {
		return true
	}
	return false
}
