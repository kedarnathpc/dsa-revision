/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func solve (root *TreeNode, sum *int) {
    if root == nil {
        return
    }

    solve (root.Right, sum)
    *sum += root.Val
    root.Val = *sum
    solve (root.Left, sum)
}

func bstToGst(root *TreeNode) *TreeNode {
    sum := 0
    solve (root, &sum)

    return root
}
