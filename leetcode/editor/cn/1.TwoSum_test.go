//Given an array of integers nums and an integer target, return indices of the
//two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may
//not use the same element twice.
//
// You can return the answer in any order.
//
//
// Example 1:
//
//
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//
//
// Example 2:
//
//
//Input: nums = [3,2,4], target = 6
//Output: [1,2]
//
//
// Example 3:
//
//
//Input: nums = [3,3], target = 6
//Output: [0,1]
//
//
//
// Constraints:
//
//
// 2 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
// Only one valid answer exists.
//
//
//
//Follow-up: Can you come up with an algorithm that is less than
//O(n²) time complexity?
//
// Related Topics 数组 哈希表 👍 15022 👎 0

package cn

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	t.Log(res)
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i := range nums {
		if v, ok := numsMap[target-nums[i]]; ok {
			return []int{v, i}
		}
		numsMap[nums[i]] = i
	}
	return []int{0, 0}
}

//leetcode submit region end(Prohibit modification and deletion)
