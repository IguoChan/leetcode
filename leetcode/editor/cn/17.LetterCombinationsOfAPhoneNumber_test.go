//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
//
// Related Topics 哈希表 字符串 回溯 👍 2167 👎 0

package cn

import (
	"testing"
)

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {
	digits := "2"
	t.Log(letterCombinations(digits))

	digits1 := "23"
	t.Log(letterCombinations(digits1))
}

//leetcode submit region begin(Prohibit modification and deletion)
var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combs []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combs = []string{} // 这一步是必须的，每次必须清空这个全局变量
	backtrack(digits, 0, "")
	return combs
}

// 注意，如果将combs传入，会无效的，除非new一个对象传入
func backtrack(digits string, idx int, comb string) {
	if idx == len(digits) {
		combs = append(combs, comb)
	} else {
		letters := phoneMap[string(digits[idx])]
		for i := range letters {
			backtrack(digits, idx+1, comb+string(letters[i]))
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
