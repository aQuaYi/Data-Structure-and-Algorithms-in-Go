package heap

import "sort"

//Interface 是排序的接口
type Interface interface {
	sort.Interface
	Divide(int, int) Interface
}

//Sort 使用堆排序
func Sort(a Interface) {
	heapify(a)
	heapSort(a)
}

//heapify 堆有序化
//heapify后，根节点是序列中的最大值。
//a[i]的son节点是a[2*i+1]和a[2*i+2]
//a[i]的father节点是a[(i-1)/2]
//root节点是a[0]
func heapify(a Interface) {
	n := a.Len()
	for i := n/2 - 1; i >= 0; i-- {
		//a[n/2-1]是序号最大的有son节点的元素
		//i:=n-1也可以堆有序化，但是，浪费了很多时间
		reheapify(a, i)
	}
}

//reheapify前，如果以i节点的两个子节点为root节点的两个子完全二叉树是堆有序的。
//reheapify后，以i节点为root节点的子完全二叉树，堆有序。而且，i节点是这个子树中的最大值。
func reheapify(a Interface, i int) {
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
		reheapify(a, 0)    //让a重新堆有序化
	}
}
