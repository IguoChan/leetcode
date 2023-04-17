//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics 数组 回溯 👍 1341 👎 0

package cn

import (
	"sort"
	"testing"
)

func TestPermutationsIi(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums) // 先给nums排个序
	ans := make([][]int, 0)

	var bt func(track []int, used []bool)
	bt = func(track []int, used []bool) {
		if len(nums) == len(track) {
			ans = append(ans, append([]int{}, track...))
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}

			// 注意这里的剪枝条件，只有当此位置和前一个位置重复，但是此位置又不是第一次的时候才能通过
			// 否则就剪枝
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			track = append(track, nums[i])
			used[i] = true

			bt(track, used)

			track = track[:len(track)-1]
			used[i] = false
		}
	}

	used := make([]bool, len(nums))
	bt([]int{}, used)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
