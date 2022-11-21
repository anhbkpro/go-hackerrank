package evaluate_boolean_binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Runtime: 11 ms, faster than 91.23% of Go online submissions for Evaluate Boolean Binary Tree.
//Memory Usage: 6.4 MB, less than 80.70% of Go online submissions for Evaluate Boolean Binary Tree.
func evaluateTree(root *TreeNode) bool {
	if root.Val <= 1 {
		return root.Val == 1
	}
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	return evaluateTree(root.Left) && evaluateTree(root.Right)
}
