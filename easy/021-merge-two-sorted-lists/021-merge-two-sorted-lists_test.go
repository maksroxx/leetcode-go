package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	list2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}
	listsResult := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{nil, list2, list2},
		{list1, nil, list1},
		{list1, list2, listsResult},
	}
	for _, tt := range tests {
		result := mergeTwoLists(tt.list1, tt.list2)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("mergeTwoLists(%+v, %+v) = %+v; want %+v", tt.list1, tt.list2, result, tt.want)
		}
	}
}
