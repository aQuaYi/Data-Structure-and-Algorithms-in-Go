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
	next       *node
	first      *node
	firstAgain bool
}

func newChain(node *node) *chain {
	if node == nil {
		return nil
	}
	return &chain{
		next:  node.next,
		first: node.next,
	}
}

func (c *chain) Next() interface{} {
	item := c.next.item
	c.next = c.next.next
	return item
}

func (c *chain) HasNext() bool {
	if c == nil { //空的链直接返回false
		return false
	}
	if c.next == c.first {
		if !c.firstAgain {
			c.firstAgain = true
			return true
		}
		return false
	}

	return true
}
