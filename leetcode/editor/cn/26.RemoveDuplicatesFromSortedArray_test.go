//给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致
//。
//
// 由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个
//元素应该保存最终结果。
//
// 将最终结果插入 nums 的前 k 个位置后返回 k 。
//
// 不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
// 判题标准:
//
// 系统会用下面的代码来测试你的题解:
//
//
//int[] nums = [...]; // 输入数组
//int[] expectedNums = [...]; // 长度正确的期望答案
//
//int k = removeDuplicates(nums); // 调用
//
//assert k == expectedNums.length;
//for (int i = 0; i < k; i++) {
//    assert nums[i] == expectedNums[i];
//}
//
// 如果所有断言都通过，那么您的题解将被 通过。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：2, nums = [1,2,_]
//解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,1,1,1,2,2,3,3,4]
//输出：5, nums = [0,1,2,3,4]
//解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 已按 升序 排列
//
//
// Related Topics 数组 双指针 👍 2907 👎 0

package cn

import (
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(removeDuplicates(nums))

	nums = []int{0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(removeDuplicates(nums))

	nums = []int{0, 1, 2}
	t.Log(removeDuplicates(nums))

	nums = []int{0, 0, 1, 2}
	t.Log(removeDuplicates(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	start, end, k := 1, 1, 1 // 双指针法，start从第二个元素开始遍历
	for start < l {
		for end < l && nums[end] <= nums[start-1] {
			// 当end所指向的值比start前一个指针所执行的值要小时，我们向前移动end
			end++
		}
		if end == l {
			break // 如果超了边界，需要返回循环
		}
		// 以下为正常情况
		nums[start], nums[end] = nums[end], nums[start] // 当end的值终于大于start-1的值时，交换end和start的值
		start++                                         // 循环条件，start后移一位
		end++                                           // 循环条件，end后移一位
		k++                                             // 累积
	}

	nums = nums[:k]
	return k
}

//leetcode submit region end(Prohibit modification and deletion)
