package main

import (
	"algo/dynamic_programming"
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(1500)
	result := dynamic_programming.Fib(n)

	fmt.Printf("result from %d is %d", n, result)
}
