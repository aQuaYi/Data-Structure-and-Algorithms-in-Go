package queue

//Iterator 是为了实现queue的迭代功能而实现的接口。
/*可以按照如下的方法，访问Iterator中的每一个元素。
for Iterator.HasNext(){
	doSomething(Iterator.Next())
}
我在学习这个的时候，比较难以想通的一点是，第一个也叫Next。
*/
type Iterator interface {
	//Next 会返回chain当前node的值
	//使用Next前，请确保HasNext()为true
	Next() interface{}

	// HasNext 检查了Next是否为nil
	HasNext() bool
}

type chain struct {
	node *node
}

func (c *chain) Next() interface{} {
	item := c.node.item
	c.node = c.node.next
	return item
}

func (c *chain) HasNext() bool {
	if c.node != nil {
		return true
	}
	return false
}
