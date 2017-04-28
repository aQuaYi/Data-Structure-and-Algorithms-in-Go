package quick

import (
	"fmt"
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

func partition(a sort.Interface) int {
	i, j := 1, a.Len()-1
	fmt.Println(a)
	for {
		for a.Less(i, 0) && i < a.Len() {
			i++
		}
		for a.Less(0, j) && j > 0 {
			j--
		}
		if i >= j {
			break
		}
		a.Swap(i, j)
		fmt.Printf("Swap(%d,%d)\n", i, j)
		fmt.Println(a)
	}
	a.Swap(0, j)
	fmt.Printf("Swap(0,%d)\n", j)
	fmt.Println(a)

	return j
}
