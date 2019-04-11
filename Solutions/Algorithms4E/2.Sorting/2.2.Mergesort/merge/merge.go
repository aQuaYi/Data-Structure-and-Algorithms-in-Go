package merge

//通过bench测试，以下merge方式最快
var merge = mergeAssignmentReslice

//对通过赋值的方式把子切片的值传递给result
//然后对子切片reslice
func mergeAssignmentReslice(left, right []int) []int {
	cap := len(left) + len(right)
	result := make([]int, cap, cap)
	for k := 0; k < cap; k++ {
		switch {
		case len(left) == 0:
			result = append(result[:k], right...)
			return result
		case len(right) == 0:
			result = append(result[:k], left...)
			return result
		case left[0] < right[0]:
			result[k] = left[0]
			left = left[1:]
		default:
			result[k] = right[0]
			right = right[1:]
		}
	}
	panic("Never Reach Here")
}

//通过赋值的方式把子切片的值传递给result
func mergeAssignment(left, right []int) []int {
	cap, i, j := len(left)+len(right), 0, 0
	result := make([]int, cap, cap)
	for k := 0; k < cap; k++ {
		switch {
		case i >= len(left):
			result[k] = right[j]
			j++
		case j >= len(right):
			result[k] = left[i]
			i++
		case left[i] < right[j]:
			result[k] = left[i]
			i++
		default:
			result[k] = right[j]
			j++
		}
	}
	return result
}

//通过append的方式把子切片的值传递给result
//传值后，对子切片reslice
func mergeReslice(left, right []int) []int {
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
