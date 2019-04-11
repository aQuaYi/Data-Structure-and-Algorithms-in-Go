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
//a.CompareTo(b) < 0  ===> a < b
//a.CompareTo(b) == 0 ===> a == b
//a.CompareTo(b) > 0  ===> a > b
type Comparer interface {
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
	//没有找到对应的key，就生成新node
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
		//找到了对应的key，就更新value
		n.value = value
	}

	//更新node的计数
	n.n = size(n.left) + size(n.right) + 1
	return n
}

//Get 获取对应key的value
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

//Floor 返回不大于key的b的最大key
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
		//key > n.key时
		t := floor(n.right, key)
		if t != nil {
			//右侧存在t.key<key时，才返回t
			return t
		}
		//否侧，返回n
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
		//key < n.key时
		t := ceiling(n.left, key)
		if t != nil {
			//左侧存在key<t.key时，返回t
			return t
		}
		//否侧，返回n
		return n
	}
}

//Select  returns BST[k]'s key
//返回某个key，BST中key小的键有k个
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

//Rank 返回比key小的键的个数。
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
	if n == nil {
		return nil
	}

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
	if n == nil {
		return nil
	}

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
		n.left = delete(n.left, key)
	case cmp > 0:
		n.right = delete(n.right, key)
	default:
		n = deleteRoot(n)
	}
	n.n = size(n.left) + size(n.right) + 1
	return n
}

func deleteRoot(n *node) *node {
	if n == nil { //如果deleteRoot只是用在delete中，可以不写这个if语句的。但是为了通用性，还是写了。
		return nil
	}

	switch {
	case n.left == nil:
		return n.right
	case n.right == nil:
		return n.left
	default:
		//删除root后，把右侧的最小值作为新的root
		t := n
		n = min(n.right)
		n.right = deleteMin(t.right)
		n.left = t.left //这一行和上一行不能互换，想想为什么？
		n.n = size(n.left) + size(n.right) + 1
		return n
	}
}
