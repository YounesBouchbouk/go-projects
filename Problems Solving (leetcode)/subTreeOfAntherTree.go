package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil {
		return false
	}

	return isSameTree(root, subRoot) || isSubtree(root.Right, subRoot) || isSubtree(root.Left, subRoot)
}
