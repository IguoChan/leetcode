//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。
//
// 注意：解集不能包含重复的组合。
//
//
//
// 示例 1:
//
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//
// 示例 2:
//
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
//
//
//
// 提示:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
//
// Related Topics 数组 回溯 👍 1314 👎 0

package cn

import (
	"sort"
	"testing"
)

func TestCombinationSumIi(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	res := combinationSum2(candidates, 8)
	t.Log(res)
}

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 首先给切片排序
	ans := make([][]int, 0)

	var bt func(target, idx int, track []int)
	bt = func(target, idx int, track []int) {
		if target == 0 {
			ans = append(ans, append([]int{}, track...))
			return
		}

		for i := idx; i < len(candidates); i++ {
			delta := target - candidates[i]
			if delta < 0 { // 因为candidates已经升序排序的原因，如果delta < 0，那么随着i增大，后续的都会<0
				break
			}

			// 这里做去重的剪枝
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			track = append(track, candidates[i])
			bt(delta, i+1, track)
			track = track[:len(track)-1]
		}
	}

	bt(target, 0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
