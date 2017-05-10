package binarySearchTree

//Node 是binary search tree所有方法的集合
type Node interface {
}

//Comparer 规定了元素的可比较性
type Comparer interface {
	//a.CompareTo(b) < 0  ===> a < b
	//a.CompareTo(b) == 0 ===> a == b
	//a.CompareTo(b) > 0  ===> a > b
	CompareTo(Comparer) int
}

type node struct {
	key         Comparer
	value       interface{}
	left, right *node
	n           int
}

//New 返回新结点
func New(key Comparer, value interface{}) Node {
	return &node{
		key:   key,
		value: value,
		n:     1,
	}
}

func (n *node) Size() int {
	if n == nil {
		return 0
	}
	return n.n
}

func (n *node) Get(key Comparer) interface{} {
	if n == nil {
		return nil
	}
	cmp := key.CompareTo(n.key)
	switch {
	case cmp < 0:
		return n.left.Get(key)
	case cmp > 0:
		return n.right.Get(key)
	default:
		return n.value
	}
}

func (n *node) Put(key Comparer, value interface{}) {
	if n == nil {
		n = New(key, value).(*node)
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp < 0:
		n.left.Put(key, value)
	case cmp > 0:
		n.right.Put(key, value)
	default:
		n.value = value
	}

	n.n = n.left.Size() + n.right.Size() + 1
}

func (n *node) MinKey() interface{} {
	return n.min().key
}

func (n *node) min() *node {
	if n == nil {
		return nil
	}
	if n.left == nil {
		return n
	}
	return n.left.min()
}

func (n *node) MaxKey() interface{} {
	return n.max().key
}

func (n *node) max() *node {
	if n == nil {
		return nil
	}
	if n.right == nil {
		return n
	}
	return n.right.max()
}

func (n *node) Floor(key Comparer) interface{} {
	if n == nil {
		return nil
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp == 0:
		return n.key
	case cmp < 0:
		return n.left.Floor(key)
	default:
		t := n.right.Floor(key)
		if t != nil {
			return t
		}
		return n.key
	}
}

func (n *node) Ceiling(key Comparer) interface{} {
	if n == nil {
		return nil
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp == 0:
		return n.key
	case cmp > 0:
		return n.right.Floor(key)
	default:
		t := n.left.Floor(key)
		if t != nil {
			return t
		}
		return n.key
	}
}

func (n *node) Select(k int) interface{} {
	if n == nil {
		return nil
	}

	t := n.Size()
	switch {
	case t > k:
		return n.left.Select(k)
	case t < k:
		return n.left.Select(k - t - 1)
	default:
		return n.key
	}
}

func (n *node) Rank(key Comparer) int {
	if n == nil {
		return 0
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp < 0:
		return n.left.Rank(key)
	case cmp > 0:
		return 1 + n.left.Size() + n.right.Rank(key)
	default:
		return n.left.Size()
	}
}

func (n *node) DeleteMin() {
	if n.left == nil {
		n = n.right
	}

	n.left.DeleteMin()
	n.n = n.left.Size() + n.right.Size() + 1
}

func (n *node) DeleteMax() {
	if n.right == nil {
		n = n.left
	}

	n.right.DeleteMax()
	n.n = n.left.Size() + n.right.Size() + 1
}

func (n *node) Delete(key Comparer) {
	if n == nil {
		return
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp < 0:
		n.left.Delete(key)
	case cmp > 0:
		n.right.Delete(key)
	default:
		n.deleteRoot()
	}

}

func (n *node) deleteRoot() {
	if n == nil {
		return
	}

	switch {
	case n.left == nil:
		n = n.right
	case n.right == nil:
		n = n.left
	default:
		t := n
		n = n.right.min()
		n.left = t.left
		t.right.DeleteMin()
		n.right = t.right
	}
	n.n = n.left.Size() + n.right.Size() + 1
}
