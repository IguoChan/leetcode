//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
//
// Related Topics 数组 二分查找 分治 👍 6006 👎 0

package cn

import (
	"testing"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	t.Log(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	t.Log(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3, 4, 9}
	nums2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(findMedianSortedArrays(nums1, nums2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	// 找出合并数组的第k个值
	var findKth func(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int
	findKth = func(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
		// 若有一个数组提前结束，则返回另一个数组的第start+k个值
		l1, l2 := end1-start1+1, end2-start2+1
		if l1 == 0 {
			return nums2[start2+k-1]
		}
		if l2 == 0 {
			return nums1[start1+k-1]
		}

		// 到最后需要找到的第一个值，那就找最小的那个
		if k == 1 {
			return min(nums1[start1], nums2[start2])
		}

		i, j := start1+min(l1, k/2)-1, start2+min(l2, k/2)-1
		if nums1[i] < nums2[j] {
			return findKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
		} else {
			return findKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
		}
	}

	n, m := len(nums1), len(nums2)
	a, b := (n+m)/2, (n+m)%2
	if b == 0 {
		return float64(findKth(nums1, 0, n-1, nums2, 0, m-1, a)+findKth(nums1, 0, n-1, nums2, 0, m-1, a+1)) / 2
	} else {
		return float64(findKth(nums1, 0, n-1, nums2, 0, m-1, a+1))
	}
}

//leetcode submit region end(Prohibit modification and deletion)
