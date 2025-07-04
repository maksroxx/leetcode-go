package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 1 2 3 4 5 nil
// 5 4 3 2 1 nil

// curr = 5
// prev = nil
// next = 4
// 4 := nil
// nil = 5
// 5 4
