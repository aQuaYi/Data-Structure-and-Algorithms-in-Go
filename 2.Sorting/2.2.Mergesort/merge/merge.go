package merge

func merge(left, right []int) []int {
	var result []int
	for {
		switch {
		case len(left) == 0:
			result = append(result, right...)
			return result
		case len(right) == 0:
			result = append(result, left...)
			return result
		case left[0] < right[0]:
			result = append(result, left[0])
			left = left[1:]
		default:
			result = append(result, right[0])
			right = right[1:]
		}
	}
}

//Sort 返回一个排序好了的int切片
func Sort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2
	return merge(Sort(a[:mid]), Sort(a[mid:]))
}

//BUSort 返回自底向上的排序好的int切片
func BUSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	n := len(a)
	var result []int
	for sz := 1; sz < n*2; sz *= 2 {
		result = make([]int, 0, n)
		for i := 0; i < n; i += sz {
			mid := (i + i + sz) / 2
			result = append(result, merge(a[i:min(mid, n)], a[min(mid, n):min(i+sz, n)])...)
		}
		a = result
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
