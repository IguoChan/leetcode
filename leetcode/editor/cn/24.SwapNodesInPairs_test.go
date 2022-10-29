//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
//
// Related Topics 递归 链表 👍 1612 👎 0

package cn

import (
	"testing"
)

func TestSwapNodesInPairs(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	ret := swapPairs(head)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}

	t.Log("-----------------")
	head = nil
	ret = swapPairs(head)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}

	t.Log("-----------------")
	head = &ListNode{
		Val:  1,
		Next: nil,
	}
	ret = swapPairs(head)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}

	t.Log("-----------------")
	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	ret = swapPairs(head)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	first, second := head, head.Next
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	last := dummy
	for second != nil {
		last.Next = second // last始终指向此次循环的上一个指针，这里将其指向second
		// 以下两步交换first和second的顺序
		first.Next = second.Next
		second.Next = first
		// 以下几步进行指针循环操作，分别是
		// last指向现有的first，也就是交换后的第二个节点，下一次交换的节点之前节点
		last = first
		// first指针指向其下一个，即下次循环的第一个节点
		first = first.Next
		// 如果first指针为空，则挑出循环，否则将second指针指向其下一个节点，即下次循环的第二个节点
		if first == nil {
			break
		}
		second = first.Next
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
