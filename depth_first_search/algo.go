package depth_first_search

import "container/list"

func DFS(root Node, val int) *Node {
	nodes := list.New()
	nodes.PushBack(&root)

	for nodes.Len() > 0 {
		node := nodes.Back().Value.(*Node)
		nodes.Remove(nodes.Back())

		if node.Val == val {
			return node
		}

		if node.Children == nil {
			continue
		}

		for _, child := range node.Children {
			nodes.PushBack(&child)
		}
	}

	return nil
}
