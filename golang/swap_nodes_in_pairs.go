/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ 
 func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    n := head.Next
    a, b := head, head.Next // next, and cur is a pointer
    prev := new(ListNode)
    for a!=nil && b != nil {
        a.Next = b.Next
        b.Next = a
        if prev != nil  {
            prev.Next = b
        }
        if a.Next == nil {
            break
        }
        b = a.Next.Next
        prev = a
        a = a.Next
    }
    return n
}

