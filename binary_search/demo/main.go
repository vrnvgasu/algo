package main

import (
	"algo/binary_search"
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		arr = append(arr, i*2)
	}

	fmt.Println("arr: ", arr)

	target := rand.Intn(20)
	fmt.Println("target: ", target)

	index := binary_search.BinarySearch(arr, target)
	fmt.Println("index: ", index)
}
