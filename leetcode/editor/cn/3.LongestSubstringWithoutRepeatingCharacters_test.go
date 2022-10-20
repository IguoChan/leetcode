//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 8289 👎 0

package cn

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	s := "abcabcbb"
	t.Log(lengthOfLongestSubstring(s))

	s = "bbbbb"
	t.Log(lengthOfLongestSubstring(s))

	s = "pwwkew"
	t.Log(lengthOfLongestSubstring(s))

	s = ""
	t.Log(lengthOfLongestSubstring(s))
}

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	repeats := make(map[uint8]int) // map表用于记录某字符最新的位置
	l := len(s)
	start, end, max := 0, 0, 0 // 双指针和最大值
	for ; end < l; end++ {
		// 首先检查s[end]这个字符是否在前出现过，且至少大于等于start，这时候将指针至少挪到i的后面
		if i, ok := repeats[s[end]]; ok && start <= i {
			start = i + 1
		}

		// 判断最大值有没有更新
		if max < end-start+1 {
			max = end - start + 1
		}

		// 将end这个字符作为key，位置作为value入表
		repeats[s[end]] = end
	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)
