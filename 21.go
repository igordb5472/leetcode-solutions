package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur := &ListNode{}
	preHead := cur
	for list1 != nil || list2 != nil {
		switch {
		case list1 == nil:
			cur.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		case list2 == nil:
			cur.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		case list1.Val <= list2.Val:
			cur.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		default:
			cur.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return preHead.Next
}
