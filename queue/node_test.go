package queue

import "testing"

func testNodeNew(t *testing.T, node *node, val interface{}) {

	if node.item != val {
		t.Fatalf("Expecting %#v but got %#v\n", val, node.item)
	}
	if node.next != nil {
		t.Fatalf("Expecting no next but got %#v\n", node.next)
	}
}

func TestNodeNew(t *testing.T) {
	node := newNode(10)
	testNodeNew(t, node, 10)
}
