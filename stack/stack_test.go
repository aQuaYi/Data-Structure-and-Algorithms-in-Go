package stack

import "testing"

var S stack

func Test_Stack_New(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Error("新生成的Stack不是空的")
	}
}

func Test_Stack_Push(t *testing.T) {
	S.Push(1)
	if S.n != 1 {
		t.Error("Stack的数量没有在Push时增加。")
	}
	if S.last.item.(int) != 1 {
		t.Error("Stack没能Push入正确的数据。")
	}
}

func Test_Stack_Pop(t *testing.T) {
	item, ok := S.Pop().(int)
	if !ok {
		t.Error("Stack没能正确地转换出节点的类型。")
	}
	if item != 1 {
		t.Error("Stack没能正确地Pop出1")
	}

	if S.n != 0 {
		t.Error("Stack在Pop后没能，个数没能减少一个。")
	}
}

func Test_Stack_Iterator(t *testing.T) {
	N := 1000

	for i := 0; i < N; i++ {
		S.Push(i)
	}

	if S.Size() != N {
		t.Errorf("Stack没能完全地Push入%d个数字", N)
	}

	c := S.Iterator()
	i := N - 1
	for c.HasNext() {
		item := c.Next().(int)
		if item != i {
			t.Error("Stack生成的迭代器不能按顺序迭代")
		}
		i--
	}
	if i != -1 {
		t.Error("Stack生成的迭代器没有完成全部迭代。")
	}

	if S.Size() != N {
		t.Error("迭代后，影响了Stack的长度。")
	}
}
