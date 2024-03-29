//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的
// 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//
//
//
// 示例 1：
//
//
//输入：candidates = [2,3,6,7], target = 7
//输出：[[2,2,3],[7]]
//解释：
//2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
//7 也是一个候选， 7 = 7 。
//仅有这两种组合。
//
// 示例 2：
//
//
//输入: candidates = [2,3,5], target = 8
//输出: [[2,2,2,2],[2,3,3],[3,5]]
//
// 示例 3：
//
//
//输入: candidates = [2], target = 1
//输出: []
//
//
//
//
// 提示：
//
//
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// candidates 的所有元素 互不相同
// 1 <= target <= 40
//
//
// Related Topics 数组 回溯 👍 2447 👎 0

package cn

import (
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 5}
	res := combinationSum(candidates, 8)
	t.Log(res)
}

// 第一种解法
func combinationSum1(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	var dfs func(target, idx int, comb []int)
	dfs = func(target, idx int, comb []int) {
		if idx == len(candidates) {
			return
		}

		if target == 0 {
			ans = append(ans, append([]int{}, comb...)) // 需要new一个slice去copy，因为后续传参会复用comb底层数组
			return
		}

		dfs(target, idx+1, comb)

		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx, comb)
		}
	}

	dfs(target, 0, []int{})

	return ans
}

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	var bt func(target, idx int, track []int)
	bt = func(target, idx int, track []int) {
		if target == 0 {
			ans = append(ans, append([]int{}, track...))
		}

		for i := idx; i < len(candidates); i++ {
			delta := target - candidates[i]
			if delta < 0 {
				continue
			}

			track = append(track, candidates[i])
			bt(delta, i, track)
			track = track[:len(track)-1]
		}
	}

	bt(target, 0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
