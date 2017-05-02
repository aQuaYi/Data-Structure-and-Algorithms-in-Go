package heap

import "sort"

//Interface 是排序的接口
type Interface interface {
	sort.Interface
	Divide(int, int) Interface
}

//Sort 使用堆排序
//a[i]的son节点是a[2*i+1]和a[2*i+2]
//a[i]的father节点是a[(i-1)/2]
//root节点是a[0]
func Sort(a Interface) {
	reheapify(a)
	heapSort(a)
}

//reheapify 采取top-down方式的堆有序化
//reheapify后，根节点是序列中的最大值。
func reheapify(a Interface) {
	n := a.Len()
	for i := (n - 2) / 2; i >= 0; i-- {
		//a[(n-2)/2]是序号最大的有son节点的元素
		//i:=n-1也可以堆有序化，但是，浪费了很多时间
		sink(a, i)
	}
}

//sink 让较小的元素下沉
func sink(a Interface, i int) {
	n := a.Len()
	for 2*i+1 < n { //当i节点存在son节点的时候，继续循环
		iSon := 2*i + 1
		if iSon+1 < n && a.Less(iSon, iSon+1) {
			//iSon+1<n 保证了a[iSon+1]不越界
			//a.Less(iSon, iSon+1)是为了i与较大的Son互换。
			iSon++
		}
		if a.Less(iSon, i) {
			return
		}
		a.Swap(i, iSon)
		i = iSon
	}
}

func heapSort(a Interface) {
	n := a.Len()
	for n > 1 {
		n--
		a.Swap(0, n)       //把序列a中的最大值，放在序列尾部
		a = a.Divide(0, n) //把已经已经归为的a的最大值与其余部分隔离开
		sink(a, 0)         //让a重新堆有序化
	}
}

func swin(a Interface, i int) {
	for i > 0 {
		iFather := (i - 1) / 2
		if !a.Less(iFather, i) {
			return
		}
		a.Swap(iFather, i)
		i = iFather
	}
}
