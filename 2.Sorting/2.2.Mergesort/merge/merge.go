package merge

func merge(left, right []int) []int {
	cap := len(left) + len(right)
	result := make([]int, 0, cap)
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
	for sz := 1; sz < n; sz *= 2 {
		result = make([]int, 0, n)
		for i := 0; i < n; i += sz + sz {
			mid := min(i+sz, n)
			hi := min(mid+sz, n)
			result = append(result, merge(a[i:mid], a[mid:hi])...)
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
