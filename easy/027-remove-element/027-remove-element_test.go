package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected int
	}{
		{
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			val:      3,
			expected: 4,
		},
		{
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
		{
			nums:     []int{},
			val:      2,
			expected: 0,
		},
	}
	for _, tt := range tests {
		result := removeElement(tt.nums, tt.val)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("removeElement(%v, %d) = %d; want %d", tt.nums, tt.val, result, tt.expected)
		}
	}
}
