//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
//
// Related Topics 链表 双指针 👍 2283 👎 0

package cn

import (
	"testing"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 哑结点，直接指向头结点
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	fast, slow := head, dummy // 快慢指针分别只想头结点和哑结点
	for i := 0; i < n; i++ {  // 因为n不会大于节点的长度，所以不会出现循环中fast为nil的情形
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 循环结束时，slow指针指向需要删除节点的前一个节点
	slow.Next = slow.Next.Next
	// 如果需要删除的正是head指针，那么在第一次循环结束时，fast就是nil了，第二次循环不会开始，此时slow指针指向dummy，也符合指向需要删除的需要删除节点的前一个节点
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
