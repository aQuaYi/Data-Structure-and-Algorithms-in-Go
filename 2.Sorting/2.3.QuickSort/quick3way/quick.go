package quick3way

import "math/rand"
import "sort"

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
	lt, gt := partition(a)
	Sort(a.Divide(0, lt))
	Sort(a.Divide(gt+1, a.Len()))
}

func partition(a Interface) (int, int) {
	lt, i, gt := 0, 1, a.Len()-1
	for i <= gt {
		switch {
		case a.Less(i, lt):
			a.Swap(lt, i)
			lt++
			i++
		case a.Less(lt, i):
			a.Swap(i, gt)
			gt--
		default:
			i++
		}
	}
	return lt, gt
}
