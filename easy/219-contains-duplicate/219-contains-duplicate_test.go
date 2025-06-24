package containsduplicate

import (
	"reflect"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for _, tt := range tests {
		result := containsNearbyDuplicate(tt.nums, tt.k)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("containsNearDuplicate(%v, %d) = %t; want %t", tt.nums, tt.k, result, tt.want)
		}
	}
}
