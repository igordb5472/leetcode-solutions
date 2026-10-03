package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists))
}

func merge(lists []*ListNode, l, r int) *ListNode {
	switch r - l {
	case 1:
		return lists[l]
	default:
		less, greater := merge(lists, l, (l+r)/2), merge(lists, (l+r)/2, r)
		if less == nil {
			return greater
		}
		if greater == nil {
			return less
		}
		if less.Val > greater.Val {
			less, greater = greater, less
		}
		headLs, headGr := less, greater
		for headLs.Next != nil {
			for headGr != nil && headGr.Val <= headLs.Next.Val {
				nextHeadGr := headGr.Next
				headGr.Next = headLs.Next
				headLs.Next = headGr
				headGr = nextHeadGr
			}
			headLs = headLs.Next
		}
		for headGr != nil {
			nextHeadGr := headGr.Next
			headGr.Next = headLs.Next
			headLs.Next = headGr
			headLs = headGr
			headGr = nextHeadGr
		}
		return less
	}
}
