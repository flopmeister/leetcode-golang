
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// recursive
func reverseList2(head *ListNode) *ListNode {
	return reverseListRecursive(nil, head)
}
func reverseListRecursive(prev *ListNode, cur *ListNode) *ListNode {
	if cur == nil {
		return prev
	}
	next := cur.Next
	cur.Next = prev
	return reverseListRecursive(cur, next)
}
