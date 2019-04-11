package  bag 

import "testing"

var emptyChain = &chain{}

func Test_chain_HasNext_emptyChain(t *testing.T) {
	if emptyChain.HasNext() {
		t.Error("empty chain have a non-nil node")
	}
}

func Test_chain_Next_HasNext(t *testing.T) {
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	one.next = two
	two.next = three
	c := &chain{node: one}
	i := 0
	for c.HasNext() {
		i++
		if c.Next().(int) != i {
			t.Error("chain.Next()没有返回正确的值")
		}
	}
	if i != 3 {
		t.Error("chain.HasNext()没能执行正确的次数。")
	}
}
