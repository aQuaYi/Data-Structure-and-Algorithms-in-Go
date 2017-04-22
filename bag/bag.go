package bag

import "sync"

//Bag 是一个具有背包api的接口。

//Bag 是线程安全的。
type Bag interface {
	//Add 往Bag内添加数据。
	Add(interface{})

	//IsEmpty 用于判断Bag是否为空
	IsEmpty() bool

	//Size 返回Bag内元素的数量。
	Size() int

	//Iterator 返回一个迭代器
	Iterator() Iterator
}

// New 返回一个Bag接口，空的哟
func New() Bag {
	return &bag{}
}

type bag struct {
	sync.RWMutex
	first *node
	n     int
}

func (b *bag) Add(item interface{}) {
	b.Lock()
	defer b.Unlock()
	newFirst := newNode(item)
	newFirst.next = b.first
	b.first = newFirst
	b.n++
}

func (b *bag) IsEmpty() bool {
	b.Lock()
	defer b.Unlock()
	return b.first == nil
}

func (b *bag) Size() int {
	b.RLock()
	defer b.RUnlock()
	return b.n
}

func (b *bag) Iterator() Iterator {
	b.RLock()
	defer b.RUnlock()
	return &chain{node: b.first}
}
