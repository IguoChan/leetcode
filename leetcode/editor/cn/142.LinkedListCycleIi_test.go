//给定一个链表的头节点 head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到
//链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
// 不允许修改 链表。
//
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
// 示例 3：
//
//
//
//
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围在范围 [0, 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
// pos 的值为 -1 或者链表中的一个有效索引
//
//
//
//
// 进阶：你是否可以使用 O(1) 空间解决此题？
//
// Related Topics 哈希表 链表 双指针 👍 2088 👎 0

package cn

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestLinkedListCycleIi(t *testing.T) {
	var second *ListNode
	second = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  -4,
				Next: second,
			},
		},
	}
	second.Next.Next.Next = second
	head := &ListNode{
		Val:  3,
		Next: second,
	}
	hasCycle := detectCycle(head)
	t.Log(hasCycle)
}

// hash表方式，空间O(N)
func detectCycle1(head *ListNode) *ListNode {
	ma := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := ma[head]; ok {
			return head
		}

		ma[head] = struct{}{}
		head = head.Next
	}

	return nil
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针方式
// 如果没有环，则返回nil，如果有环：
// 设环前一段长度为a，环起点到相遇点距离为b，相遇点再到环起点距离为c
// 设fast指针到相遇点跑过的距离为f，slow指针的距离为s
// 则有 f = a + n * (b+c) + b，（n >= 1，因为fast很快，如果n=0，则速度和slow一样了）
// 可以证明，在s在环内的一圈内必定和fast相遇，因为在s刚到环起点时，fast已经在环内的某点了，而fast
// 的速度是slow的两点倍，假设在slow进入环起点的时候，二者差距x，0<=x<(b+c)；slow的速度为1/step,
// fast的速度为2/step，假设在n step后二者相遇，则 2n-n=x，即n=x,所以0<=n<(b+c)，所以相遇必定在slow运行一周内
// 所以 s = a + b
// 而有f = 2*s， 则a + n * (b+c) + b = 2a + 2b，得到 a = (n-1)(b+c) + c
// 以上公式说明，a的长度等于圈长度的整数倍+c

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 此时证明肯定有环，然后根据公式 a = (n-1)(b+c) + c,
			// 使用fast充当新的指针从head继续轮询，直到和slow再次相遇，则相遇点肯定是环起点
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
