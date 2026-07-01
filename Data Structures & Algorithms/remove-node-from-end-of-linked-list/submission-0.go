/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // need to know the position of the n as well as the max length
	// store the list to array/slices
	arrNode := make([]*ListNode,0)
	currHead := head
	for currHead != nil {
		arrNode = append(arrNode, currHead)
		currHead = currHead.Next
	}

	// Get the position of the node in the slices len(array) - n
	// main logic:
	deletedPos := len(arrNode)-n
	// what if on the edge / first node -> prevent index out of bound -> return pos.Next
	if deletedPos == 0 {
		return arrNode[0].Next
	} 
	// pos-1.Next = pos.Next -> return idx 0
	curr := arrNode[deletedPos]
	arrNode[deletedPos-1].Next = curr.Next
	return arrNode[0]
}
