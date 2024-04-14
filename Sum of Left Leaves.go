/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode) int {
    if root == nil {
        return 0
    }

    ans := 0
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        ans += root.Left.Val
    }

    ans += traverse(root.Left)
    ans += traverse(root.Right)

    return ans
}


func sumOfLeftLeaves(root *TreeNode) int {
    return traverse(root)
}s
