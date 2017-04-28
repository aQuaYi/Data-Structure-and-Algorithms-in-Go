package quick

import (
	"math/rand"
	"sort"
)

//Interface 是quick排序的接口
type Interface interface {
	sort.Interface
	Divide(int, int) Interface
}

//Sort 使用quickSort算法进行排序。
func Sort(a Interface) {
	if a.Len() <= 1 {
		return
	}

	j := rand.Intn(a.Len())
	a.Swap(0, j)
	j = partition(a)
	Sort(a.Divide(0, j))
	Sort(a.Divide(j+1, a.Len()))
}

//partition 要考虑好，当a.Len()==2时，如何排好序
func partition(a sort.Interface) int {
	i, j := 1, a.Len()-1
	for {
		for !a.Less(0, i) && i < a.Len()-1 {
			// i<a.Len()-1 保证了当i不会为a.Len()
			// 例如，当a[0]为最大值时，i会一直变大。直到i++后，i==a.Len()
			// 那么在下一个a.Less(i,0)时，会产生"index out of range"错误。
			// 当a.Len()==2时，这种情况很容易发生。
			i++
		}
		for !a.Less(j, 0) && j > 0 {
			// j>0, 保证了j不会为-1
			// 例如，当a[0]为最小值时，j会一直变小。直到j--后，j==-1
			// 那么在下一个a.Less(0,j)时，会产生"index out of range"错误。
			// 当a.Len()==2时，这种情况很容易发生。
			j--
		}
		if i >= j {
			break
		}
		a.Swap(i, j)
	}
	a.Swap(0, j)
	return j
}
