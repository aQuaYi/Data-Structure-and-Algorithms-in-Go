package shell

import (
	"sort"
)

//Sort 对a进行希尔排序
func Sort(a sort.Interface) {
	n := a.Len()
	h := 1
	factor := 3
	for h < n/factor {
		h = factor*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && a.Less(j, j-h); j -= h {
				a.Swap(j, j-h)
			}
		}
		h /= factor
	}
}
