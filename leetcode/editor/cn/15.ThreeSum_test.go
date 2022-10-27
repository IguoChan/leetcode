//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
//k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
// 你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//
//
// 示例 3：
//
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics 数组 双指针 排序 👍 5340 👎 0

package cn

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(nums))

	nums = []int{0, 1, 1}
	t.Log(threeSum(nums))

	nums = []int{0, 0, 0}
	t.Log(threeSum(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	for first := 0; first < l; first++ {
		// 第一个数字的枚举值需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := l - 1
		for second := first + 1; second < l; second++ {
			// 第二个数字的枚举值也需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 如果三者相加大于0，那么一直递减third
			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}

			// 如果第二个数和第三个数已经重合，那么后续就不会有满足的数据了，因为nums已经被排序
			if second == third {
				break
			}

			// 匹配条件
			if nums[first]+nums[second]+nums[third] == 0 {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
