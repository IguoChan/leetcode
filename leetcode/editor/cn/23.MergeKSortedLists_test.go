//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2228 👎 0

package cn

import (
	"math"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestMergeKSortedLists(t *testing.T) {
	//lists = [[1,4,5],[1,3,4],[2,6]]
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	lists := []*ListNode{l1, l2, l3}
	l := mergeKLists(lists)
	for l != nil {
		t.Log(l.Val)
		l = l.Next
	}

	t.Log("---------------------------------")

	lists = nil
	l = mergeKLists(lists)
	for l != nil {
		t.Log(l.Val)
		l = l.Next
	}

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// 找到lists中每个链表第一个节点中最小节点
	findMin := func(lists []*ListNode) (int, *ListNode) {
		minVal := math.MaxInt32
		var minList *ListNode
		idx := 0
		for i, l := range lists {
			ll := l
			if ll != nil && ll.Val < minVal {
				minList = ll
				idx = i
				minVal = ll.Val
			}
		}
		return idx, minList
	}

	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	op := dummy
	for {
		idx, minList := findMin(lists)
		if minList == nil {
			break
		}
		op.Next = minList
		op = op.Next
		minList = minList.Next
		lists[idx] = minList
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
