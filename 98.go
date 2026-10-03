package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, -(1<<31)-1, (1 << 31))
}

func validate(root *TreeNode, minLeftVal, maxRightVal int) bool {
	validRootVal := root.Val > minLeftVal && root.Val < maxRightVal
	validLeftSubtree := root.Left == nil || root.Left.Val < root.Val && validate(root.Left, minLeftVal, root.Val)
	validRightSubtree := root.Right == nil || root.Right.Val > root.Val && validate(root.Right, root.Val, maxRightVal)
	return validRootVal && validLeftSubtree && validRightSubtree
}
