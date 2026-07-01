/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // loop through nodes till empty
	// put into array
	arr := make([]*ListNode,0)
	currHead := head
	for currHead != nil {
		arr = append(arr,currHead)
		currHead=currHead.Next
	}

	// now we know start and finish
	// using 2 points at the start and end of the array
	// while start < end
	idxA := 0
	idxB := len(arr)-1
	for idxA < idxB && idxA+1!=idxB{
	// redirect the next
	// start -> end
	// end-1 -> null
	// end-> start+1
		arr[idxA].Next = arr[idxB]
		arr[idxB-1].Next = arr[idxB].Next
		arr[idxB].Next = arr[idxA+1]
		idxA++
		idxB--
	}
}
