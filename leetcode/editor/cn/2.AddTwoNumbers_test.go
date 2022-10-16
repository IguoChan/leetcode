//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
//
// Related Topics 递归 链表 数学 👍 8767 👎 0

package cn

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{}
	//s1 := []int{2, 4, 3}
	//s1 := []int{0}
	s1 := []int{9, 9, 9, 9, 9, 9, 9}
	next := l1
	var pre *ListNode
	for _, s := range s1 {
		pre = next
		pre.Val = s
		next = &ListNode{}
		pre.Next = next
	}
	pre.Next = nil

	l2 := &ListNode{}
	//s2 := []int{5, 6, 4}
	//s2 := []int{0}
	s2 := []int{9, 9, 9, 9}
	next = l2
	for _, s := range s2 {
		pre = next
		pre.Val = s
		next = &ListNode{}
		pre.Next = next
	}
	pre.Next = nil

	res := addTwoNumbers(l1, l2)
	var s3 []int
	for res != nil {
		s3 = append(s3, res.Val)
		res = res.Next
	}

	t.Log(s3)
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode // 表示最后返回链表的头尾指针
	carry := 0               // 进位数

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			// 第一次
			head = &ListNode{Val: sum}
			tail = head
		} else {
			// 之后的循环，tail表示此次循环中上一节点
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	// 如果有进位，需要加上
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
