package queue

import "testing"

var Q queue

func Test_Queue_new(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Error("新生成的Queue不是空的")
	}
}

func Test_Queue_Enqueue(t *testing.T) {
	Q.Enqueue(1)
	if Q.first != Q.last {
		t.Fatal("向空Queue中Enqueue第一个元素时，Q.first和Q.last没有指向同一个node。")
	}
	if Q.n != 1 {
		t.Error("Queue的个数没有在Enqueue中增加。")
	}
}

func Test_Queue_Dequeue(t *testing.T) {
	item, ok := Q.Dequeue().(int)
	if !ok {
		t.Error("queue没能正确地转换出queue中节点的类型。")
	}
	if item != 1 {
		t.Error("queue没能正确地Dequeue出1")
	}

	if Q.n != 0 {
		t.Error("queue在dequeue后，queue的数量没有减少一个。")
	}
}

func Test_Queue_Iterator(t *testing.T) {
	N := 1000

	for i := 0; i < N; i++ {
		Q.Enqueue(i)
	}

	if Q.Size() != N {
		t.Errorf("queue没能完全地Enqueue入%d个数字", N)
	}

	c := Q.Iterator()
	i := 0
	for c.HasNext() {
		item := c.Next().(int)
		if item != i {
			t.Error("queue生成的迭代器不能按顺序迭代")
		}
		i++
	}
	if i != N {
		t.Error("queue生成的迭代器没有完成全部迭代。")
	}

	if Q.Size() != N {
		t.Error("迭代后，影响了Q的长度。")
	}
}
