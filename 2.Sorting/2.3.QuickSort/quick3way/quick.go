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
		return //长度为1或者空序列，无需排序。
	}

	//随机选取一个数字，与列首的元素进行交换，可以提高快排的性能
	j := rand.Intn(a.Len())
	a.Swap(0, j)

	lt, gt := partition(a)

	Sort(a.Divide(0, lt))
	Sort(a.Divide(gt+1, a.Len()))
}

//partition把序列划分成三段
//假设排序前，a[0]==v
//那么排序后，
//a[:lt]中的元素都<v
//a[lt,gt+1]中的所有元素都是v
//a[gt+1:]中的元素都>v
func partition(a Interface) (int, int) {
	lt, i, gt := 0, 1, a.Len()-1
	for i <= gt { //i>gt说明，
		//无论lt为何值，a[lt]始终为v
		switch {
		case a.Less(i, lt): //把比a[lt]小的值，移动到lt前面去。
			a.Swap(lt, i) //这个步骤确保了，a[:lt]中的元素都<v
			lt++
			i++
		case a.Less(lt, i): //把比a[lt]大的值，移动到尾部去。
			a.Swap(i, gt) //这个步骤确保了，a[gt+1:]中都是>v的元素。
			gt--
			//这里没有i++是因为，这个新被移动到i位的元素，还没有与v比较过，需要在下一循环中与a[lt]比较。
		default: //a[i]==a[lt]
			i++
		}
	}
	return lt, gt
}
