/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil  {
		return nil
	}
	
	if root.Val < key {
		root.Right = deleteNode(root.Right,key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left,key)
	} else { // if match
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minimumNode := minNode(root.Right)
			root.Val=minimumNode.Val
			root.Right = deleteNode(root.Right,minimumNode.Val)
		}

	}

	return root
}

func minNode(root *TreeNode) *TreeNode{
	for root.Left != nil {
		root=root.Left
	}
	return root
}
