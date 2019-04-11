package bag

import "testing"

var B bag

func Test_Bag_New(t *testing.T) {
	b := New()
	if !b.IsEmpty() {
		t.Error("新生成的Bag不是空的")
	}
}

func Test_Bag_Add(t *testing.T) {
	B.Add(0)
	if B.n != 1 {
		t.Error("Bag的数量没有在Add时增加。")
	}
	if B.last.item.(int) != 0 {
		t.Error("Bag没能Add入正确的数据。")
	}
}

func Test_Bag_Iterator(t *testing.T) {
	N := 1000

	for i := 1; i < N; i++ {
		B.Add(i)
	}

	if B.Size() != N {
		t.Errorf("Bag没能完全地Add入%d个数字", N)
	}

	c := B.Iterator()
	i := N - 1
	for c.HasNext() {
		item := c.Next().(int)
		if item != i {
			t.Error("Bag生成的迭代器不能按顺序迭代")
		}
		i--
	}
	if i != -1 {
		t.Error("Bag生成的迭代器没有完成全部迭代。")
	}

	if B.Size() != N {
		t.Error("迭代后，影响了Bag的长度。")
	}
}
