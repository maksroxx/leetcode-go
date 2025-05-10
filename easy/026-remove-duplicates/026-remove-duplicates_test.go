package removeduplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 2, 3}, 3},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, tt := range tests {
		result := removeDuplicates(tt.nums)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("removeDuplicates(%v) = %d; want %d", tt.nums, result, tt.expected)
		}
	}
}
