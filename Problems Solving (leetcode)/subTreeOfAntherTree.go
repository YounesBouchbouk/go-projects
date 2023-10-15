package main

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil {
		return false
	}

	return isSameTree(root, subRoot) || IsSubtree(root.Right, subRoot) || IsSubtree(root.Left, subRoot)
}
