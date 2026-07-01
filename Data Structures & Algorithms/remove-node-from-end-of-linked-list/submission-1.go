/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // find total length iterate the node through
	totalLength := 0
	currHead := head
	for currHead != nil {
		currHead =currHead.Next
		totalLength++
	}
	// find deleted position -> total - n 
	deletedPosition := totalLength - n
	// if 0 -> return head.Next
	if deletedPosition==0{
		return head.Next
	}

	// else
	// traverse from head 
	traversePosition := 0
	curr := head
	for curr != nil {
		// if position == deletedPosition-1
		if traversePosition==deletedPosition-1{

			// position.Next = position.Next.Next
			// return head
			curr.Next=curr.Next.Next
			return head
		}
		curr = curr.Next
		traversePosition++
	}

	return nil
}
