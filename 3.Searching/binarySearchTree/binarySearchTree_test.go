package binarySearchTree

import "testing"
import "fmt"

type intComparable int

func (ic intComparable) CompareTo(b Comparer) int {
	return int(ic - b.(intComparable))
}

func (n *node) equals(a *node) bool {
	if n == nil && a == nil {
		return true
	}

	if n.Size() != a.Size() {
		return false
	}

	if n.key != a.key || n.value != a.value {
		return false
	}

	return n.left.equals(a.left) && n.right.equals(a.right)
}
func Test_Node_equals(t *testing.T) {
	var testNew *node
	if !testNew.equals(nil) {
		t.Error("equals认为两个nil不相等")
	}

	testNew = New(intComparable(5), 0).(*node)
	testB := &node{
		key:   intComparable(5),
		value: 0,
		n:     1,
	}
	if !testNew.equals(testB) {
		t.Error("equals不能发现相通的node")
	}

	testB = &node{
		key:   intComparable(4), //原key为5
		value: 0,
		n:     1,
	}
	if testNew.equals(testB) {
		t.Error("equals无法发现不同的key")
	}

	testB = &node{
		key:   intComparable(5),
		value: 1, //原value为0
		n:     1,
	}
	if testNew.equals(testB) {
		t.Error("equals无法发现不同的value")
	}

	testB = &node{
		key:   intComparable(5),
		value: 0,
		n:     2, //原n为1
	}
	if testNew.equals(testB) {
		t.Error("equals无法发现不同的n")
	}
}

func Test_Node_New(t *testing.T) {
	testNew := New(intComparable(5), 0).(*node)
	testB := &node{
		key:   intComparable(5),
		value: 0,
		n:     1,
	}
	if !testNew.equals(testB) {
		t.Error("New新生成的node不正确")
	}
}

func Test_Node_Size(t *testing.T) {
	var tns *node
	if tns.Size() != 0 {
		t.Error("nil的size!=0")
	}

	tns = &node{
		key:   intComparable(5),
		value: 0,
		n:     1,
	}
	if tns.Size() != 1 {
		t.Error("无法正确读取*node的Size")
	}
}

func Test_Node_Put(t *testing.T) {
	ics := []intComparable{5, 3, 7, 1, 4, 6, 9, 2, 8}

	var tns *node
	for i := 0; i < len(ics); i++ {
		fmt.Printf("tns's Pointer=%p,\n", tns)
		tns.Put(ics[i], i)

		if tns.Size() != i+1 {
			t.Errorf("在put时，不能正确地修改Size， tns.Size()=%d, 实际Size=%d", tns.Size(), i+1)
		}
	}
}
