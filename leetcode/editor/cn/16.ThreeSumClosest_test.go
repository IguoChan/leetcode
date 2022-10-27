//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 双指针 排序 👍 1273 👎 0

package cn

import (
	"math"
	"sort"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 100, 1, -4}
	t.Log(threeSumClosest(nums, 1))

	nums = []int{0, 0, 3}
	t.Log(threeSumClosest(nums, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	if l <= 3 {
		return nums[0] + nums[1] + nums[2]
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	minDelta := math.MaxInt32 // 表示sum-target
	for first := 0; first < l; first++ {
		third := l - 1 // 使用双指针，第三个数从排序后数组最大值开始
		for second := first + 1; second < third; {
			delta := nums[first] + nums[second] + nums[third] - target
			if delta == 0 {
				return target
			}
			if abs(delta) < abs(minDelta) {
				minDelta = delta
			}

			if nums[first]+nums[second]+nums[third] < target {
				second++
			} else {
				third--
			}
		}
	}

	return minDelta + target
}

//leetcode submit region end(Prohibit modification and deletion)
