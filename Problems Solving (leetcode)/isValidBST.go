package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(node *TreeNode, lower *int, upper *int) bool {
	if node == nil {
		return true
	}

	val := node.Val
	if lower != nil && val <= *lower {
		return false
	}
	if upper != nil && val >= *upper {
		return false
	}

	if !helper(node.Right, &val, upper) {
		return false
	}
	if !helper(node.Left, lower, &val) {
		return false
	}

	return true
}
