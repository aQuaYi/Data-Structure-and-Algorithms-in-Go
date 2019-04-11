package bag 

type node struct {
	item interface{}
	next *node
}

func newNode(item interface{}) *node {
	return &node{
		item: item,
	}
}
