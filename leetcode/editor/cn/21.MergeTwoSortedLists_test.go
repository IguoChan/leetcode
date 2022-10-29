//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
//
// Related Topics 递归 链表 👍 2760 👎 0

package cn

import (
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{ // 哑结点，用作定位头结点
		Val:  0,
		Next: nil,
	}
	l := dummy // 循环节点
	for list1 != nil || list2 != nil {
		if list1 == nil { // 如果有链表为空了，那么直接将另一个链表挂在l节点的下一个节点上即可，并挑出循环
			l.Next = list2
			break
		}
		if list2 == nil { // 同上
			l.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			l.Next = list1     // 确定l所指向节点的下一节点是较小值
			l = l.Next         // 循环条件：将l指向我们需要返回链表的最后一个节点
			list1 = list1.Next // 循环条件：将list1指针指向其下一个节点
		} else { // 同上
			l.Next = list2
			l = l.Next
			list2 = list2.Next
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
