package merge_sort

func MergeSort(list []int) []int {
	if len(list) == 1 {
		return list
	}
	left, right := []int{}, []int{}
	length := len(list)
	center := length / 2
	for i := 0; i < length; i++ {
		if i < center {
			left = append(left, list[i])
		} else {
			right = append(right, list[i])
		}
	}

	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left []int, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	length := len(left) + len(right)
	result := make([]int, 0, length)

	l, r := 0, 0
	for i := 0; i < length; i++ {
		if leftLen-1 < l {
			result = append(result, right[r])
			r++

			continue
		}
		if rightLen-1 < r {
			result = append(result, left[l])
			l++

			continue
		}

		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	return result
}
