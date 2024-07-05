package main

import (
	dfs "algo/depth_first_search"
	"fmt"
	"math/rand"
)

func main() {
	root := dfs.Node{
		Val: rand.Intn(10),
		Children: []dfs.Node{
			{
				Val: rand.Intn(10),
				Children: []dfs.Node{
					{Val: rand.Intn(10)},
					{Val: rand.Intn(10)},
					{Val: rand.Intn(10)},
				},
			},
			{
				Val: rand.Intn(10),
				Children: []dfs.Node{
					{Val: rand.Intn(10)},
					{Val: rand.Intn(10)},
					{Val: rand.Intn(10)},
				},
			},
			{
				Val: rand.Intn(10),
				Children: []dfs.Node{
					{Val: rand.Intn(10)},
					{Val: rand.Intn(10)},
					{Val: rand.Intn(10)},
					{Val: -1},
				},
			},
		},
	}

	fmt.Println("graph: ", root)

	val := rand.Intn(10)
	fmt.Println("val: ", val)

	node := dfs.DFS(root, val)
	if node == nil {
		fmt.Println("not found")
		return
	}
	fmt.Println("node: ", *node)

}
