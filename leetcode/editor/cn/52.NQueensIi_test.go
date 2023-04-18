//n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 回溯 👍 423 👎 0

package cn

import (
	"bytes"
	"testing"
)

func TestNQueensIi(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	res := 0

	valid := func(row, col int, track [][]byte) bool {
		n := len(track)
		// 检查列是否有皇后冲突
		for i := 0; i < n; i++ {
			if track[i][col] == 'Q' {
				return false
			}
		}
		// 检查右上方是否有皇后冲突
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if track[i][j] == 'Q' {
				return false
			}
		}
		// 检查左上方是否有皇后冲突
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if track[i][j] == 'Q' {
				return false
			}
		}
		return true
	}

	var bt func(idx int, track [][]byte)
	bt = func(row int, track [][]byte) {
		if row == n {
			res++
			return
		}

		for col := 0; col < n; col++ {
			if !valid(row, col, track) {
				continue
			}

			track[row][col] = 'Q'
			bt(row+1, track)
			track[row][col] = '.'
		}

	}

	track := make([][]byte, n)
	for i := range track {
		track[i] = bytes.Repeat([]byte{'.'}, n)
	}
	bt(0, track)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
