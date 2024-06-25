package main

import (
	"algo/merge_sort"
	"fmt"
	"math/rand"
)

func main() {
	const size = 20
	list := make([]int, 0, size)
	for i := 0; i < size; i++ {
		v := rand.Intn(100)
		if rand.Int()%2 == 0 {
			v *= -1
		}
		list = append(list, v)
	}
	fmt.Println("list: ", list)

	sortedList := merge_sort.MergeSort(list)
	fmt.Println("sortedList: ", sortedList)
}
