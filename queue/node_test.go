package queue

import "testing"

func testStateNew(t *testing.T, node *node, val interface{}) {

	if node.item != val {
		t.Fatalf("Expecting %#v but got %#v\n", val, node.item)
	}
	if node.next != nil {
		t.Fatalf("Expecting no next but got %#v\n", node.next)
	}
}

func TestStateNew(t *testing.T) {
	state := newState(10)
	testStateNew(t, state, 10)
}
