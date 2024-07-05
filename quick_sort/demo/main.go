package main

import (
	"algo/quick_sort"
	"fmt"
	"math/rand"
)

const items = 10

func main() {
	arr := make([]int, 0, items)
	for i := 0; i < items; i++ {
		arr = append(arr, rand.Intn(100))
	}

	fmt.Println("not sorted:", arr)

	arr = quick_sort.QuickSort(arr)
	fmt.Println("sorted:", arr)
}
