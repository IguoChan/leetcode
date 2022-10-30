//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
//
//提示：
//
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//
//
//
// Related Topics 递归 链表 👍 1839 👎 0

package cn

import (
	"testing"
)

func TestReverseNodesInKGroup(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 标准的链表反转函数
	reverse := func(head *ListNode) (pre *ListNode) {
		for head != nil {
			next := head.Next // 记录循环节点的下一个节点
			head.Next = pre   // 将循环节点的下一个节点指向前一个节点，第一次时前一个节点是空
			pre = head        // 循环条件，将下一个节点指向当前节点
			head = next       // 循环条件，将当前节点指向下一个节点
		}
		return pre
	}

	// 链表题会发生头结点变更的，最好都设置一个哑结点用于记录头结点变更
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	pre, end := dummy, dummy
	for end.Next != nil { // end指针循环前会始终指向下一次要开始循环的节点的前一个节点，所以终止条件是其后还有节点
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break // 当最后剩下的待翻转节点数小于k时，无需反转，所以这里直接挑出循环
		}
		start, next := pre.Next, end.Next // start这一小段反转后的最后一个节点，end记录翻转前end指向的下一节点，用作之后搭桥
		end.Next = nil                    // 截断需要翻转的这一小段链表
		pre.Next = reverse(start)         // 反转这一小段，并将pre的后一个节点设置为反转后的头节点
		start.Next = next                 // 搭桥，将这一小段反转后，再次接上
		pre, end = start, start           // 循环条件，pre和end都指向下一次循环开始节点的前一个节点
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
