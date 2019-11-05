package leetcode203

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		delNode := head
		head = head.Next
		delNode.Next = nil

	}
	if head == nil {
		return nil
	}

	prev := head
	// 这时候prev后面的节点才是要删除的元素
	for prev.Next != nil {
		if prev.Next.Val == val {
			delNode := prev.Next
			prev.Next = delNode.Next
			delNode.Next = nil
		} else {
			prev = prev.Next
		}
	}

	return head
}
