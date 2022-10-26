//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 5844 👎 0

package cn

import (
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	s := "babad"
	t.Log(longestPalindrome(s))

	s = "aaaa"
	t.Log(longestPalindrome(s))

	s = "cbbd"
	t.Log(longestPalindrome(s))
}

// 从中心拓展的方式
func longestPalindrome1(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}

	// 从中心拓展
	maxL, maxR := 0, 0
	for i := 0; i < l-1; i++ {
		// 首先判断是否是偶数的回文字符串
		tmpL1, tmpR1 := i, i
		if s[i] == s[i+1] {
			tmpR1++
			for offset := 1; i-offset >= 0 && i+offset+1 < l; offset++ {
				if s[i-offset] == s[i+1+offset] {
					tmpL1--
					tmpR1++
				} else {
					break
				}
			}
		}

		//其次判断是够是奇数回文字符串
		tmpL2, tmpR2 := i, i
		for offset := 1; i-offset >= 0 && i+offset < l; offset++ {
			if s[i-offset] == s[i+offset] {
				tmpL2--
				tmpR2++
			} else {
				break
			}
		}

		if tmpR1-tmpL1 > maxR-maxL {
			maxL, maxR = tmpL1, tmpR1
		}
		if tmpR2-tmpL2 > maxR-maxL {
			maxL, maxR = tmpL2, tmpR2
		}

		// 很重要的一步，判断如果以第i个数或者i/i+1为中心的回文字符串的最大长度是否已经
		// 小于现有的最大长度，如果是的话，那继续循环下去没有意义，直接退出循环即可
		if l-i <= (maxR-maxL)/2 {
			break
		}
	}

	return s[maxL : maxR+1]
}

// 动态规划的方式
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}

	dp := make([][]bool, l)
	for i := range s {
		dp[i] = make([]bool, l)
		dp[i][i] = true // 边界条件，每个字符都是回文子串
	}

	start, max := 0, 1
	// 循环顺序很重要，必须保证dp[i+1][j-1]比dp[i][j]先轮询到，所以我们以j从小到大开始轮询
	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if i+1 == j {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if max < j-i+1 && dp[i][j] {
				max = j - i + 1
				start = i
			}
		}
	}

	// 以下这种轮询会导致dp[i+1][j-1]比dp[i][j]后轮询到，导致bug
	//for i := 0; i < l-1; i++ {
	//	for j := i + 1; j < l; j++ {
	//		if s[i] != s[j] {
	//			dp[i][j] = false
	//		} else {
	//			if i+1 == j {
	//				dp[i][j] = true
	//			} else {
	//				dp[i][j] = dp[i+1][j-1]
	//			}
	//		}
	//
	//		if max < j-i+1 && dp[i][j] {
	//			max = j - i + 1
	//			start = i
	//		}
	//	}
	//}

	return s[start : start+max]
}

//leetcode submit region end(Prohibit modification and deletion)
