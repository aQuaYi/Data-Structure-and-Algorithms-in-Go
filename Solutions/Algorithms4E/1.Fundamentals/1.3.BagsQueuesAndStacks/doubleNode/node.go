package doubleNode

type node struct {
	item  interface{}
	first *node
	last  *node
}

func newNode(item interface{}) *node {
	return &node{
		item: item,
	}
}

func (n *node) Equal(item interface{}) bool {
	if n.item == item {
		return true
	}

	return false
}
