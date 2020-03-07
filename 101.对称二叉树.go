/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	return a(root,root)

}
func a(n1,n2 *TreeNode) bool{

	if n1==nil && n2==nil{
		return true
	}
	if n1==nil || n2==nil{
		return false
	}
	return n1.Val==n2.Val && a(n1.Left,n2.Right) && a(n1.Right,n2.Left)
}
// @lc code=end

