package subarraysum

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		numbers []int
		k       int
		want    int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, 2, 1, 2, 1}, 3, 4},
		{[]int{1, -1, 0}, 0, 3},
		{[]int{0, 0, 0}, 0, 6},
		{[]int{1, 2, 3, 4, 5}, 15, 1},
		{[]int{1, 2, 3}, 6, 1},
		{[]int{}, 0, 0},
	}
	for _, tt := range tests {
		result := subarraySum(tt.numbers, tt.k)
		if result != tt.want {
			t.Errorf("subarraySum(%v, %d) = %d; want %d", tt.numbers, tt.k, result, tt.want)
		}
	}
}
