/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func calculateSum(root *TreeNode, currSum int) int {
    if root == nil {
        return 0
    }

    currSum = currSum * 10 + root.Val

    if root.Left == nil && root.Right == nil {
        return currSum
    }

    return calculateSum(root.Left, currSum) + calculateSum(root.Right, currSum)
}

func sumNumbers(root *TreeNode) int {
    return calculateSum(root, 0)    
}
