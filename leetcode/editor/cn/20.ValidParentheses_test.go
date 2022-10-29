//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
//
// Related Topics 栈 字符串 👍 3578 👎 0

package cn

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	s := "([)]"
	t.Log(isValid(s))

	s = "{[()]}"
	t.Log(isValid(s))

	s = "()"
	t.Log(isValid(s))

	s = "()[]{}"
	t.Log(isValid(s))

	s = "()["
	t.Log(isValid(s))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	l := len(s)
	if l%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := range s {
		if pairs[s[i]] > 0 { // 说明此时往栈中push一个右括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				// 如果是入栈的第一个字符，或者其与栈中前一个字符不是匹配字符，则说明不可能匹配成功
				return false
			}
			// 否则缩短栈长度
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	// 如果最后栈不为空，那么说明匹配失败
	if len(stack) > 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
