package binary_search

func BinarySearch(arr []int, target int) int {
	return binarySearch(arr, target, 0)
}

func binarySearch(arr []int, target int, offset int) int {
	if len(arr) == 0 {
		return -1
	}

	i := len(arr) / 2

	if arr[i] == target {
		return i + offset
	}

	if target > arr[i] {
		if i+1 == len(arr) { // справа больше нет значений
			return -1
		}

		return binarySearch(arr[i:], target, i+offset)
	}

	return binarySearch(arr[:i], target, offset)
}
