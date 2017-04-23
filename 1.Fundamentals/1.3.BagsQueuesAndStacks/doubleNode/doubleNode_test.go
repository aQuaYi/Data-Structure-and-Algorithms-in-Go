package doubleNode

import (
	"testing"
)

var c = &chain{}

func isEquals(a, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

func Test_chain_AddFirst(t *testing.T) {
	c.AddFirst(0)
	if !isEquals(c.first.item, c.last.item) ||
		!isEquals(c.first.item, 0) {
		t.Error("往空chain中添加数据时，出错")
	}
	c = &chain{}
	n := 10
	for i := 0; i < n; i++ {
		c.AddFirst(i)
	}
	l := c.last
	for i := 0; i < n; i++ {
		if !isEquals(l.item, i) {
			t.Error("AddFirst出错")
		}
		l = l.first
	}

	f := c.first
	for i := n - 1; i >= 0; i-- {
		if !isEquals(f.item, i) {
			t.Error("AddFirst出错")
		}
		f = f.last
	}
}

func Test_chain_AddLast(t *testing.T) {
	c = &chain{}
	c.AddLast(0)
	if !isEquals(c.first.item, c.last.item) ||
		!isEquals(c.first.item, 0) {
		t.Error("往空chain中添加数据时，出错")
	}
	c = &chain{}
	n := 10
	for i := 0; i < n; i++ {
		c.AddLast(i)
	}
	l := c.last
	for i := n - 1; i >= 0; i-- {
		if !isEquals(l.item, i) {
			t.Error("AddFirst出错")
		}
		l = l.first
	}

	f := c.first
	for i := 0; i < n; i++ {
		if !isEquals(f.item, i) {
			t.Error("AddFirst出错")
		}
		f = f.last
	}
}

func Test_chain_DeleteFirst(t *testing.T) {
	c = &chain{}
	n := 10
	for i := 0; i < n; i++ {
		c.AddLast(i)
	}
	for i := 0; i < n; i++ {
		if !isEquals(c.DeleteFirst(), i) {
			t.Error("DeleteFirst出错。")
		}
	}
}

func Test_chain_DeleteLast(t *testing.T) {
	c = &chain{}
	n := 10
	for i := 0; i < n; i++ {
		c.AddFirst(i)
	}
	for i := 0; i < n; i++ {
		if !isEquals(c.DeleteLast(), i) {
			t.Error("DeleteFirst出错。")
		}
	}
}
