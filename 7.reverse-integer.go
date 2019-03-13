// package leetcode

import "math"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.18%)
 * Total Accepted:    625.8K
 * Total Submissions: 2.5M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */
func reverse(x int) int {
	if !(math.MinInt32 < x && math.MaxInt32 > x) {
		return 0
	}
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	if math.MinInt32 < rev && math.MaxInt32 > rev {
		return rev
	} else {
		return 0
	}

}
