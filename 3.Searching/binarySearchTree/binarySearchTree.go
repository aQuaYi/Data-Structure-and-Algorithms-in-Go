package binarySearchTree

//BST 是binary search tree所有方法的集合
type BST interface {
	Size() int
	Put(Comparer, interface{})
	Get(Comparer) interface{}
	Delete(Comparer)
	DeleteMax()
	DeleteMin()
	Rank(Comparer) int
	Select(int) Comparer
	Floor(Comparer) Comparer
	Ceiling(Comparer) Comparer
	Max() Comparer
	Min() Comparer
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
	n           int
	left, right *node
}

type binarySearchTree struct {
	root *node
}

func newNode(key Comparer, value interface{}) *node {
	return &node{
		key:   key,
		value: value,
		n:     1,
	}
}

//New 返回符合BST接口的数据
func New() BST {
	return &binarySearchTree{
		root: nil,
	}
}

func (b *binarySearchTree) Size() int {
	return size(b.root)
}

func size(n *node) int {
	if n == nil {
		return 0
	}
	return n.n
}

func (b *binarySearchTree) Put(key Comparer, value interface{}) {
	b.root = put(b.root, key, value)
}

func put(n *node, key Comparer, value interface{}) *node {
	if n == nil {
		return newNode(key, value)
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp < 0:
		n.left = put(n.left, key, value)
	case cmp > 0:
		n.right = put(n.right, key, value)
	default:
		n.value = value
	}
	n.n = size(n.left) + size(n.right) + 1
	return n
}

func (b *binarySearchTree) Get(key Comparer) interface{} {
	return get(b.root, key)
}

func get(n *node, key Comparer) interface{} {
	if n == nil {
		return nil
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp < 0:
		return get(n.left, key)
	case cmp > 0:
		return get(n.right, key)
	default:
		return n.value
	}
}

//Min retruns the minimum key of binary search tree
func (b *binarySearchTree) Min() Comparer {
	x := min(b.root)
	if x == nil {
		return nil
	}
	return x.key
}

func min(n *node) *node {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return n
	}

	return min(n.left)
}

//Max returns the maximum key of binary search tree
func (b *binarySearchTree) Max() Comparer {
	x := max(b.root)
	if x == nil {
		return nil
	}
	return x.key
}

func max(n *node) *node {
	if n == nil {
		return nil
	}

	if n.right == nil {
		return n
	}

	return max(n.right)
}

func (b *binarySearchTree) Floor(key Comparer) Comparer {
	x := floor(b.root, key)
	if x == nil {
		return nil
	}

	return x.key
}

func floor(n *node, key Comparer) *node {
	if n == nil {
		return nil
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp == 0:
		return n
	case cmp < 0:
		return floor(n.left, key)
	default:
		t := floor(n.right, key)
		if t != nil {
			return t
		}
		return n
	}
}

func (b *binarySearchTree) Ceiling(key Comparer) Comparer {
	x := ceiling(b.root, key)
	if x == nil {
		return nil
	}

	return x.key
}

func ceiling(n *node, key Comparer) *node {
	if n == nil {
		return nil
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp == 0:
		return n
	case cmp > 0:
		return ceiling(n.right, key)
	default:
		t := ceiling(n.left, key)
		if t != nil {
			return t
		}
		return n
	}
}

//Select  returns [k]'s key
func (b *binarySearchTree) Select(k int) Comparer {
	x := selecting(b.root, k)
	if x == nil {
		return nil
	}
	return x.key
}

//Return Node containing key of rank k
func selecting(n *node, k int) *node {
	if n == nil {
		return nil
	}

	t := size(n.left)
	switch {
	case t > k:
		return selecting(n.left, k)
	case t < k:
		return selecting(n.right, k-t-1)
	default:
		return n
	}
}

func (b *binarySearchTree) Rank(key Comparer) int {
	return rank(key, b.root)
}

func rank(key Comparer, n *node) int {
	if n == nil {
		return 0
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp < 0:
		return rank(key, n.left)
	case cmp > 0:
		return 1 + size(n.left) + rank(key, n.right)
	default:
		return size(n.left)
	}
}

func (b *binarySearchTree) DeleteMin() {
	b.root = deleteMin(b.root)
}

func deleteMin(n *node) *node {
	if n.left == nil {
		return n.right
	}

	n.left = deleteMin(n.left)
	n.n = size(n.left) + size(n.right) + 1
	return n
}

func (b *binarySearchTree) DeleteMax() {
	b.root = deleteMax(b.root)
}

func deleteMax(n *node) *node {
	if n.right == nil {
		return n.left
	}

	n.right = deleteMax(n.right)
	n.n = size(n.left) + size(n.right) + 1
	return n
}

func (b *binarySearchTree) Delete(key Comparer) {
	b.root = delete(b.root, key)
}

func delete(n *node, key Comparer) *node {
	if n == nil {
		return nil
	}

	cmp := key.CompareTo(n.key)
	switch {
	case cmp < 0:
		return delete(n.left, key)
	case cmp > 0:
		return delete(n.right, key)
	default:
		return deleteRoot(n)
	}
}

func deleteRoot(n *node) *node {
	if n == nil {
		return nil
	}

	switch {
	case n.left == nil:
		return n.right
	case n.right == nil:
		return n.left
	default:
		t := n
		n = min(n.right)
		n.left = t.left
		n.right = deleteMin(t.right)
	}

	n.n = size(n.left) + size(n.right) + 1
	return n
}
