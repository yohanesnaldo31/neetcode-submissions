/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}

	lastHead := head
	if head.Next != nil {
		lastHead = reverseList(head.Next)
		head.Next.Next = head
	}
	head.Next=nil
	return lastHead
}
