/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {  
    invertTreeChild(root)
    return root
}


func invertTreeChild(root *TreeNode) {
    if root == nil {
        return
    }

    tempVar := root.Left
    root.Left = root.Right
    root.Right= tempVar
    invertTreeChild(root.Left)
    invertTreeChild(root.Right)
}
