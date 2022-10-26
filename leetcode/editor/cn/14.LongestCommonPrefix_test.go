//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics 字符串 👍 2522 👎 0

package cn

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return strs[0]
	}

	i := 0
	for i = range strs[0] {
		for j := 1; j < l; j++ {
			// 如果任意比较子串的长度已经低于现在长度，则返回前子串
			// 如果任意比较子串的字符不等于扫描到的字符，则返回前子串
			if len(strs[j]) < i+1 || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:len(strs[0])]
}

//leetcode submit region end(Prohibit modification and deletion)
