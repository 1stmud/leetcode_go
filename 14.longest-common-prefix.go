/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (33.07%)
 * Total Accepted:    419.7K
 * Total Submissions: 1.3M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	b := false
	var str string
	for {
		var code string
		var waitCompare string
		for k, word := range strs {
			if k == 0 {
				if len(word) <= i {
					b = true
					break
				}
				waitCompare = string(word[i])
			} else {
				if len(word) <= i {
					b = true
					break
				}
				code = string(word[i])
				if code != waitCompare {
					b = true
					break
				}
			}
		}
		if b {
			break
		}
		waitCompare = code
		str += waitCompare
		i++
	}
	return str

}

