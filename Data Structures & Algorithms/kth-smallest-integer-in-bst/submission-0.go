/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int,0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr,root.Val)
		dfs(root.Right)
	}

	dfs(root)

	return arr[k-1]
}

