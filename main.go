package main

func main() {
	a := []int{1, 2, 3, 4, 5}
	longestConsecutive(a)
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

type ListNode struct {
	Val  int
	Next *ListNode
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
