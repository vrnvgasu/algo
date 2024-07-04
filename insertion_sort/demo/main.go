package main

import (
	"algo/insertion_sort"
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	fmt.Println("not sorted:", arr)

	arr = insertion_sort.InsertionSort(arr)
	fmt.Println("sorted:", arr)
}
