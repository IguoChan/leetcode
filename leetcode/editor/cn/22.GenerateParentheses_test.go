//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 👍 2927 👎 0

package cn

import (
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	t.Log(generateParenthesis(3))
}

var ss []string

// 此方法比较差
func generateParenthesis1(n int) []string {
	pair := "()"

	var f func(offset, d, n int, s string)
	f = func(offset, d, n int, s string) {
		s = s[:offset] + pair + s[offset:]
		if d == n {
			ss = append(ss, s)
			return
		}
		for i := len(s) / 2; i >= 0; i-- {
			f(i, d+1, n, s)
		}
	}

	ss = []string{}
	f(0, 1, n, "")

	res := make([]string, 0, len(ss))
	m := make(map[string]struct{})
	for _, s := range ss {
		if _, ok := m[s]; ok {
			continue
		}
		m[s] = struct{}{}
		res = append(res, s)
	}
	return res
}

//leetcode submit region begin(Prohibit modification and deletion)
/*
1、初始时定义序列的左括号数量lc 和右括号数量rc都为0。
2、如果 lc < n，左括号的个数小于n，则在当前序列str后拼接左括号。
3、如果 rc < n && lc > rc , 右括号的个数小于左括号的个数，则在当前序列str后拼接右括号。
4、当lc == n && rc == n 时，将当前合法序列str加入答案数组res中。

作者：林小鹿
链接：https://leetcode.cn/problems/generate-parentheses/solutions/938191/shen-du-you-xian-bian-li-zui-jian-jie-yi-ypti/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	dfs(n, 0, 0, "")
	return res
}

func dfs(n, lc, rc int, path string) {
	if n == lc && n == rc {
		res = append(res, path)
	} else {
		if lc < n {
			dfs(n, lc+1, rc, path+"(")
		}
		if rc < lc {
			dfs(n, lc, rc+1, path+")")
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
