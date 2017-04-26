package shell

import "testing"

type shell []int

func (s shell) Len() int {
	return len(s)
}

func (s shell) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s shell) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s shell) IsSorted() bool {
	for i := 1; i < s.Len(); i++ {
		if s.Less(i, i-1) {
			return false
		}
	}
	return true
}

func Test_Sort(t *testing.T) {
	s := make(shell, 10)
	for i := 0; i < s.Len(); i++ {
		s[i] = s.Len() - i - 1
	}
	t.Log(s)
	Sort(s)
	t.Log(s)
	if !s.IsSorted() {
		t.Error("没能排序好")
	}
}
