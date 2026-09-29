package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev, next *ListNode
	for head.Next != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	head.Next = prev
	return head
}
