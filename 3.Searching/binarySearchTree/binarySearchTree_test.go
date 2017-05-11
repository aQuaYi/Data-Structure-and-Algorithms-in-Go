package binarySearchTree

import "testing"

type intComparable int

func (ic intComparable) CompareTo(b Comparer) int {
	return int(ic - b.(intComparable))
}

func (n *node) equals(a *node) bool {
	if n == nil && a == nil {
		return true
	}

	if size(n) != size(a) {
		return false
	}

	if n.key != a.key || n.value != a.value {
		return false
	}

	return n.left.equals(a.left) && n.right.equals(a.right)
}
func Test_node_equals(t *testing.T) {
	var testA *node
	if !testA.equals(nil) {
		t.Error("equals认为两个nil不相等")
	}

	testA = &node{
		key:   intComparable(5),
		value: 0,
		n:     1,
	}
	testB := &node{
		key:   intComparable(5),
		value: 0,
		n:     1,
	}
	if !testA.equals(testB) {
		t.Error("equals不能发现相等的node")
	}

	testB = &node{
		key:   intComparable(4), //原key为5
		value: 0,
		n:     1,
	}
	if testA.equals(testB) {
		t.Error("equals无法发现不同的key")
	}

	testB = &node{
		key:   intComparable(5),
		value: 1, //原value为0
		n:     1,
	}
	if testA.equals(testB) {
		t.Error("equals无法发现不同的value")
	}

	testB = &node{
		key:   intComparable(5),
		value: 0,
		n:     2, //原n为1
	}
	if testA.equals(testB) {
		t.Error("equals无法发现不同的n")
	}
}

func Test_newNode(t *testing.T) {
	testNew := newNode(intComparable(5), 0)
	testB := &node{
		key:   intComparable(5),
		value: 0,
		n:     1,
	}
	if !testNew.equals(testB) {
		t.Error("newNode新生成的node不正确")
	}
}

func Test_New(t *testing.T) {
	tn := New().(*binarySearchTree)
	if tn.root != nil {
		t.Error("New生成的BST的root不为nil")
	}
}

func Test_size_node(t *testing.T) {
	var tns *node
	if size(tns) != 0 {
		t.Error("nil的size!=0")
	}

	tns = &node{
		key:   intComparable(5),
		value: 0,
		n:     1,
	}
	if size(tns) != 1 {
		t.Error("无法正确读取*node的Size")
	}
}

func Test_binarySearchTree_Put(t *testing.T) {
	testB := &node{
		key:   intComparable(5),
		value: 0,
		n:     1,
	}

	tb := New()
	tb.Put(intComparable(5), 0)

	if !testB.equals(tb.(*binarySearchTree).root) {
		t.Error("*binarySearchTree.Put修改binarySearchTree.root的值。")
	}
}

func Test_put(t *testing.T) {
	//验证n
	ics := []intComparable{5, 3, 7, 1, 4, 6, 9, 8, 2}
	tb := New()
	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
		if tb.Size() != i+1 {
			t.Errorf("在put时，不能正确地修改Size， tns.Size()=%d, 实际Size=%d", tb.Size(), i+1)
		}
	}

	//验证put
	ics = []intComparable{5, 3, 7, 7}
	tb = New()
	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}

	test5 := &node{key: intComparable(5), value: 0, n: 1}
	test3 := &node{key: intComparable(3), value: 1, n: 1}
	test7 := &node{
		key:   intComparable(7),
		value: 3, //最后一次发送7的value为3
		n:     1,
	}
	test5.left = test3
	test5.right = test7
	test5.n = 3

	tbRoot := tb.(*binarySearchTree).root
	if !tbRoot.equals(test5) {
		t.Error("")
	}
}

func Test_Get(t *testing.T) {
	ics := []intComparable{5, 3, 7, 1, 4, 6, 9, 8, 2}
	tb := New()
	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}

	for i := 0; i < len(ics); i++ {
		g := tb.Get(intComparable(ics[i]))
		if g.(int) != i {
			t.Error("无法Get到正确的value")
		}
	}
}

func Test_Min(t *testing.T) {
	ics := []intComparable{5, 3, 7, 1, 4, 6, 9, 8, 2}
	tb := New()

	if tb.Min() != nil {
		t.Error("nil的min key不是nil")
	}

	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}
	if tb.Min() != intComparable(1) {
		t.Error("Min无法找到正确的最小值")
	}
}

func Test_Max(t *testing.T) {
	ics := []intComparable{5, 3, 7, 1, 4, 6, 9, 8, 2}
	tb := New()

	if tb.Max() != nil {
		t.Error("nil的max key不是nil")
	}

	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}
	if tb.Max() != intComparable(9) {
		t.Error("Min无法找到正确的最大值")
	}
}

func Test_Floor(t *testing.T) {
	ics := []intComparable{5, 2, 1, 4, 7, 6, 9} //without 3,8
	iAsk := []intComparable{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iAns := []intComparable{1, 2, 2, 4, 5, 6, 7, 7, 9}

	tb := New()
	if tb.Floor(intComparable(3)) != nil {
		t.Error("nil.Floor != nil")
	}

	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}

	for i, v := range iAsk {
		if tb.Floor(v) != iAns[i] {
			t.Errorf("Floor没能返回正确的值,Floor(%d)=%d", v, tb.Floor(v))
		}
	}
}

func Test_Ceiling(t *testing.T) {
	ics := []intComparable{5, 2, 1, 4, 7, 6, 9} //without 3,8
	iAsk := []intComparable{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iAns := []intComparable{1, 2, 4, 4, 5, 6, 7, 9, 9}

	tb := New()
	if tb.Ceiling(intComparable(3)) != nil {
		t.Error("nil.Ceiling != nil")
	}

	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}

	for i, v := range iAsk {
		if tb.Ceiling(v) != iAns[i] {
			t.Errorf("Ceiling没能返回正确的值,Ceiling(%d)=%d", v, tb.Ceiling(v))
		}
	}
}

func Test_Select(t *testing.T) {
	ics := []intComparable{5, 2, 1, 4, 3, 7, 6, 8, 9}
	iAns := []intComparable{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tb := New()
	if tb.Select(3) != nil {
		t.Error("nil.Select(3) != nil")
	}

	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}
	for i, v := range iAns {
		if tb.Select(i) != v {
			t.Errorf("Select没能返回正确的值,Select(%d)=%d", i, tb.Select(i))
		}
	}
}

func Test_Rank(t *testing.T) {
	ics := []intComparable{5, 2, 1, 4, 3, 7, 6, 8, 9}
	iAns := []intComparable{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tb := New()
	if tb.Rank(intComparable(3)) != 0 {
		t.Error("nil.Rank(3) != 0")
	}

	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}

	for i, v := range iAns {
		if tb.Rank(v) != i {
			t.Errorf("Rank没能返回正确的值,Rank(%d)=%d", v, tb.Rank(v))
		}
	}

	if tb.Rank(intComparable(10)) != 9 {
		t.Error("tb.Rank(10)!=9")
	}

	if tb.Rank(intComparable(0)) != 0 {
		t.Error("tb.Rank(0)!=0")
	}
}

func Test_DeletMin(t *testing.T) {
	ics := []intComparable{5, 3, 7}
	tb := New()
	for i := 0; i < len(ics); i++ {
		tb.Put(ics[i], i)
	}

	test5 := &node{key: intComparable(5), value: 0, n: 1}
	test7 := &node{
		key:   intComparable(7),
		value: 2, 
		n:     1,
	}
	test5.right = test7
	test5.n = 2

	tbRoot := tb.(*binarySearchTree).root
	if !tbRoot.equals(test5) {
		t.Error("")
	}
}
