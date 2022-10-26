//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
//
//
//
// 示例 1：
//
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
//
//
//输入：height = [1,1]
//输出：1
//
//
//
//
// 提示：
//
//
// n == height.length
// 2 <= n <= 10⁵
// 0 <= height[i] <= 10⁴
//
//
// Related Topics 贪心 数组 双指针 👍 3863 👎 0

package cn

import (
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	start, end := 0, len(height)-1
	maxArea := 0
	for start < end {
		l := end - start
		h := 0
		if height[start] > height[end] {
			h = height[end]
			end-- // 向内移动较短的那根板
		} else {
			h = height[start]
			start++ // 向内移动较短的那根板
		}
		area := l * h
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}

//leetcode submit region end(Prohibit modification and deletion)
