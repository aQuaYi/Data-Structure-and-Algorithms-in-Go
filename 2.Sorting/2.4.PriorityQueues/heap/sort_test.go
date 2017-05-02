package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type heap []int

func (h heap) Len() int {
	return len(h)
}

func (h heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heap) Divide(i, j int) Interface {
	return h[i:j]
}

func (h heap) IsSorted() bool {
	for i := 1; i < h.Len(); i++ {
		if h.Less(i, i-1) {
			return false
		}
	}
	return true
}

func Test_Sort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	h := make(heap, 11)
	for i := 0; i < len(h); i++ {
		h[i] = len(h) - i - 1
	}
	fmt.Println(h)
	Sort(h)
	fmt.Println(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}

func Test_Sort_Sorted(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	h := make(heap, 11)
	for i := 0; i < len(h); i++ {
		h[i] = i
	}
	fmt.Println(h)
	Sort(h)
	fmt.Println(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}

func Test_Sort_with_Repeating(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := 10
	h := make(heap, n*3)
	for i := 0; i < n; i++ {
		h[i] = n - i - 1
		h[i+n], h[i+n+n] = h[i], h[i]
	}
	fmt.Println(h)
	Sort(h)
	fmt.Println(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}

func Test_Sort_100(t *testing.T) {
	h := make(heap, 100)
	for i := 0; i < h.Len(); i++ {
		h[i] = h.Len() - i - 1
	}
	Sort(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}

func Test_Sort_1000(t *testing.T) {
	h := make(heap, 1000)
	for i := 0; i < h.Len(); i++ {
		h[i] = h.Len() - i - 1
	}
	Sort(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}

}

func Test_Sort_10000(t *testing.T) {
	h := make(heap, 10000)
	for i := 0; i < h.Len(); i++ {
		h[i] = h.Len() - i - 1
	}
	Sort(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}

func Test_Sort_100000(t *testing.T) {
	h := make(heap, 100000)
	for i := 0; i < h.Len(); i++ {
		h[i] = h.Len() - i - 1
	}
	Sort(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}

func Test_Sort_1000000(t *testing.T) {
	h := make(heap, 1000000)
	for i := 0; i < h.Len(); i++ {
		h[i] = h.Len() - i - 1
	}
	Sort(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}

func Test_Sort_10000000(t *testing.T) {
	h := make(heap, 10000000)
	for i := 0; i < h.Len(); i++ {
		h[i] = h.Len() - i - 1
	}
	Sort(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}

func Test_Sort_100000000(t *testing.T) {
	h := make(heap, 100000000)
	for i := 0; i < h.Len(); i++ {
		h[i] = h.Len() - i - 1
	}
	Sort(h)
	if !h.IsSorted() {
		t.Error("没能排序好")
	}
}
