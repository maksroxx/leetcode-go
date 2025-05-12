package majorityelement

import (
	"reflect"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{2, 2, 2, 3, 3, 3, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, 5},
	}
	for _, tt := range tests {
		result := majorityElement(tt.nums)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("majorityElement(%v) = %d; want %d", tt.nums, result, tt.expected)
		}
	}
}
