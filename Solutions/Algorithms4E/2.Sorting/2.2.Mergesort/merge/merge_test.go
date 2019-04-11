package merge

import "testing"
import "math/rand"
import "time"
import "sort"

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

func Test_Sort(t *testing.T) {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = len(s) - i - 1
	}
	t.Log(s)
	s = Sort(s)
	t.Log(s)
	if !isSorted(s) {
		t.Error("没能排序好")
	}
}

func Test_Sort_1000000(t *testing.T) {
	s := make([]int, 1000000)
	for i := 0; i < len(s); i++ {
		s[i] = len(s) - i - 1
	}
	s = Sort(s)
	if !isSorted(s) {
		t.Error("没能排序好")
	}
}

func Test_BUSort(t *testing.T) {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = len(s) - i - 1
	}
	t.Log(s)
	s = BUSort(s)
	t.Log(s)
	if !isSorted(s) {
		t.Error("没能排序好")
	}
}

func Test_BUSort_1000000(t *testing.T) {
	s := make([]int, 1000000)
	for i := 0; i < len(s); i++ {
		s[i] = len(s) - i - 1
	}
	s = BUSort(s)
	if !isSorted(s) {
		t.Error("没能排序好")
	}
}

func makeIntSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	len := rand.Intn(n) + n
	is := make([]int, len, len)
	for i := 0; i < len; i++ {
		is[i] = rand.Intn(n * n)
	}
	sort.Ints(is)
	return is
}

const (
	N = 1000
)

var l, r []int

func Benchmark_MergeAssignment(b *testing.B) {
	l, r = makeIntSlice(N), makeIntSlice(N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bl, br := make([]int, len(l)), make([]int, len(r))
		copy(bl, l)
		copy(br, r)
		b.StartTimer()
		mergeAssignment(bl, br)
	}
}

func Benchmark_MergeAssignmentReslice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bl, br := make([]int, len(l)), make([]int, len(r))
		copy(bl, l)
		copy(br, r)
		b.StartTimer()
		mergeAssignmentReslice(bl, br)
	}
}
func Benchmark_MergeReslice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bl, br := make([]int, len(l)), make([]int, len(r))
		copy(bl, l)
		copy(br, r)
		b.StartTimer()
		mergeReslice(bl, br)
	}
}
