//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
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
// Related Topics 数组 回溯 👍 1731 👎 0

package cn

import (
	"bytes"
	"reflect"
	"testing"
	"unsafe"
)

func TestNQueens(t *testing.T) {
	solveNQueens(4)
}

//leetcode submit region begin(Prohibit modification and deletion)

func safeStringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var b []byte
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pbytes.Data = sh.Data
	pbytes.Len = sh.Len
	pbytes.Cap = sh.Len

	return b
}

func safeBytesTostring(bs []byte) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))

	var s string
	ps := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	ps.Data = sh.Data
	ps.Len = sh.Len
	ps.Cap = sh.Len

	return s
}

func solveNQueens(n int) [][]string {
	ans := make([][]string, 0, n)

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
			var tmp []string
			for i := range track {
				tmp = append(tmp, string(track[i]))
			}
			ans = append(ans, tmp)
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

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
