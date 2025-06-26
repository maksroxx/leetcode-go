package linkedlistcycle

import (
	"reflect"
	"testing"
)

func TestLinkedListCycle(t *testing.T) {
	// first
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	// second
	node5 := &ListNode{Val: 1}
	node6 := &ListNode{Val: 2}
	node5.Next = node6

	tests := []struct {
		head *ListNode
		want bool
	}{
		{node1, true},
		{node5, false},
		{nil, false},
	}
	for _, tt := range tests {
		result := hasCycle(tt.head)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("hasCycle(%+v) = %t; want %t", tt.head, result, tt.want)
		}
	}
}
