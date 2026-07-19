/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // create a new list node called result, and have a pointer called curr pointing to it, and next is nil
    // loop throught the linked list starting from the head
    // while l1 or l2 not nil, get the value if not nil and combine it
    // if the combine value is equal or more than 10, store 1 value, and reduce the combine value by 10
    // we create a new node which is filled with the value
    // use the stored value for the next iteration
    // if stored value still exists and loop ended, add a new node val 1 and put it to the next result
    // return result.Next

    result := &ListNode{
        Val:0,
    }
    curr := result

    var leading1 bool

    for l1 != nil || l2 != nil {
        var valueL1, valueL2 int
        if l1 != nil {
            valueL1 = l1.Val
            l1=l1.Next
        }
        if l2 != nil {
            valueL2 = l2.Val
            l2=l2.Next
        }
        var leftOver int
        if leading1{
            leftOver++
            leading1=false
        }
        combinedValue := valueL1+valueL2+leftOver
        if combinedValue >= 10 {
            leading1=true
            combinedValue=combinedValue-10
        }
        newNode := &ListNode{
            Val: combinedValue,
        }
        curr.Next=newNode
        curr=curr.Next
    }

    if leading1{
        newNode := &ListNode{
            Val: 1,
        }
        curr.Next=newNode
        curr=curr.Next
    }

    return result.Next
}