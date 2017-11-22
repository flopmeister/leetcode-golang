/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    size := 1
    cur := head
    for cur.Next != nil {
        size ++
        cur = cur.Next
    }
    fmt.Println(size)
    tail := cur
    steps := size - k%size
    for i:=0; i<steps;i++{
        tail.Next = head
        tail = head
        head = head.Next
        tail.Next = nil
    }
    return head
}