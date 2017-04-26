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
