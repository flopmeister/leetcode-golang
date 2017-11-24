/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
	vals := []int{}
	head := ListNode{}
	ptr := &head
	for i := 0; i < len(lists); i++ {
		for lists[i] != nil {
			vals = append(vals, lists[i].Val)
			lists[i] = lists[i].Next
		}
	}
	sort.Ints(vals)
	for _, val := range vals {
		ptr.Next = &ListNode{Val: val}
		ptr = ptr.Next
	}
	return head.Next
}

// func mergeKLists(lists []*ListNode) *ListNode {
//     res := ListNode{}
//     if len(lists) == 0{
//         return res.Next
//     }
//     lp := []*ListNode{}
//     for i:= range lists{
//         if lists[i] != nil {
//             lp = append(lp, lists[i])
//         }
//     }
//     head := &res
//     for len(lp)>0{
//         cur := 0
//         for i:=1;i<len(lp);i++{
//             if lp[i]!= nil &&lp[i].Val < lp[cur].Val {
//                 cur = i
//             }
//         }
//         head.Next = lp[cur]
//         head = head.Next
//         if lp[cur].Next == nil {
//             if cur == len(lp) - 1 {
//                 lp = lp[0:cur]
//             }else{
//                 lp = append(lp[0:cur], lp[cur+1:]...)
//             }
//         }else{
//             lp[cur] = lp[cur].Next
//         }
//     }
//     return res.Next
// }
