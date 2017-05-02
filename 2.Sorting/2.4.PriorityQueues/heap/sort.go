package heap

import (
	"sort"
)

//Sort 使用堆排序
func Sort(a sort.Interface) {
	n := a.Len()
	for i := (n - 1) / 2; i >= 0; i-- {
		sink(a, i)
	}
	for n > 0 {
		a.Swap(1, n)
		n--
		sink(a, 1)
	}
}

func swin(a sort.Interface, i int) {
	for i > 0 {
		iFather := (i - 1) / 2
		if !a.Less(iFather, i) {
			return
		}
		a.Swap(iFather, i)
		i = iFather
	}
}

func sink(a sort.Interface, i int) {
	n := a.Len()
	for 2*i+1 < n {
		iSon := 2*i + 1
		if iSon+1 < n && a.Less(iSon, iSon+1) {
			//iSon+1<n 是为了确保a[iSon+1]不越界
			//a.Less(iSon, iSon+1)是为了与较大的Son互换。
			iSon++
		}
		if !a.Less(i, iSon) {
			return
		}
		a.Swap(i, iSon)
		i = iSon
	}
}
