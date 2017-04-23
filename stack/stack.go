package stack

import "sync"

//Stack 是一个具有栈api的接口。
//线程安全
type Stack interface {
	//Push 往栈内添加数据。
	Push(interface{})

	//Pop 弹出最后放入栈内的数据。
	Pop() interface{}

	//IsEmpty 检查栈是否为空
	IsEmpty() bool

	//Size 返回栈内元素的数量。
	Size() int

	//Iterator 返回一个迭代器
	Iterator() Iterator
}

// New 返回一个队列接口，空的哟
func New() Stack {
	return &stack{}
}

type stack struct {
	sync.RWMutex
	last *node //指向最新的node
	n    int   //记录长度
}

func (s *stack) Push(item interface{}) {
	s.Lock()
	defer s.Unlock()
	oldLast := s.last
	s.last = newNode(item)
	s.last.next = oldLast
	s.n++
}

func (s *stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	s.Lock()
	defer s.Unlock()
	item := s.last.item
	s.last = s.last.next
	s.n--
	return item
}

func (s *stack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.last == nil
}

func (s *stack) Size() int {
	s.RLock()
	defer s.RUnlock()
	return s.n
}

func (s *stack) Iterator() Iterator {
	s.RLock()
	defer s.RUnlock()
	return &chain{node: s.last}
}
