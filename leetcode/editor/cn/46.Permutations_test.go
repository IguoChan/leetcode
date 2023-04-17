//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 2516 👎 0

package cn

import (
	"testing"
)

func TestPermutations(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	ans := make([][]int, 0)

	var bt func(track []int, used []bool)
	bt = func(track []int, used []bool) {
		if len(track) == len(nums) {
			ans = append(ans, append([]int{}, track...)) // track底层是一个数组，所以每次append需要new一份新的切片
			return
		}

		for i := range nums {
			if used[i] { // 剪枝
				continue
			}
			// 做选择
			track = append(track, nums[i])
			used[i] = true
			// 进入下一层决策树
			bt(track, used)
			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	used := make([]bool, len(nums)) // used在bt函数中直接寻址，所以需要new一下，指定len
	bt([]int{}, used)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
