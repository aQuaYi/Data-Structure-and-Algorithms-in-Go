package bag

import "sync"

//Bag 是一个具有背包api的接口。
//线程安全
type Bag interface {
	//Add 往栈内添加数据。
	Add(interface{})

	//IsEmpty 检查背包是否为空
	IsEmpty() bool

	//Size 返回背包内元素的数量。
	Size() int

	//Iterator 返回一个迭代器
	Iterator() Iterator
}

// New 返回一个队列接口，空的哟
func New() Bag {
	return &bag{}
}

type bag struct {
	sync.RWMutex
	last *node //指向最新的node
	n    int   //记录长度
}

func (b *bag) Add(item interface{}) {
	b.Lock()
	defer b.Unlock()
	oldLast := b.last
	b.last = newNode(item)
	b.last.next = oldLast
	b.n++
}

func (b *bag) IsEmpty() bool {
	b.RLock()
	defer b.RUnlock()
	return b.last == nil
}

func (b *bag) Size() int {
	b.RLock()
	defer b.RUnlock()
	return b.n
}

func (b *bag) Iterator() Iterator {
	b.RLock()
	defer b.RUnlock()
	return &chain{node: b.last}
}
