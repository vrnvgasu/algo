package quick_sort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	midIndex := len(arr) / 2
	leftItems := make([]int, 0)
	rightItems := make([]int, 0)

	for i, _ := range arr {
		if i == midIndex {
			continue
		}

		if arr[i] < arr[midIndex] {
			leftItems = append(leftItems, arr[i])
		} else {
			rightItems = append(rightItems, arr[i])
		}
	}

	leftItems = QuickSort(leftItems)
	rightItems = QuickSort(rightItems)

	res := make([]int, 0, len(arr))
	res = append(res, leftItems...)
	res = append(res, arr[midIndex])
	res = append(res, rightItems...)

	return res
}
