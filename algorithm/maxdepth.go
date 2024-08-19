package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodes := []*TreeNode{root}

	maxDeep := 0

	for len(nodes) > 0 {
		maxDeep++

		size := len(nodes)

		for i := 0; i < size; i++ {
			currentNode := nodes[0]

			nodes = nodes[1:]

			if currentNode.Left != nil {
				nodes = append(nodes, currentNode.Left)
			}

			if currentNode.Right != nil {
				nodes = append(nodes, currentNode.Right)
			}
		}
	}

	return maxDeep
}
