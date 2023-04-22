//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
//
// 返回符合要求的 最少分割次数 。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：1
//解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：0
//
//
// 示例 3：
//
//
//输入：s = "ab"
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 661 👎 0

package cn

import (
	"math"
	"testing"
)

func TestPalindromePartitioningIi(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func minCut(s string) int {
	min := math.MaxInt

	isValid := func(s string) bool {
		l, r := 0, len(s)-1
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}

		return true
	}

	var bt func(idx int, track []string)
	bt = func(idx int, track []string) {
		if idx == len(s) {
			if len(track)-1 < min {
				min = len(track) - 1
			}
			return
		}

		for i := idx; i < len(s); i++ {
			str := s[idx : i+1]
			if !isValid(str) {
				continue
			}

			track = append(track, str)
			bt(idx+len(str), track)
			track = track[:len(track)-1]
		}
	}

	bt(0, []string{})

	return min
}

//leetcode submit region end(Prohibit modification and deletion)
