package doubleNode

import "testing"

func testNodeNew(t *testing.T, node *node, val interface{}) {
	if node.item != val {
		t.Fatalf("Expecting %#v but got %#v\n", val, node.item)
	}
	if node.first != nil && node.last != nil {
		t.Fatalf("Expecting no next but got %#v, %#v\n", node.first, node.last)
	}
}

func Test_node_New(t *testing.T) {
	node := newNode(10)
	testNodeNew(t, node, 10)
}

func Test_node_Equal(t *testing.T) {
	n := newNode(1)
	if !n.Equal(1) {
		t.Error("*node.Equal()把相等的判断为不相等。")
	}
	if n.Equal(0) {
		t.Error("*node.Equal()把不相等的判断为相等。")
	}
}
