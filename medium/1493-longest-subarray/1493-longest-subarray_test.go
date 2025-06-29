package longestsubarray

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
		{[]int{}, 0},
	}
	for _, tt := range tests {
		result := longestSubarray(tt.nums)
		if result != tt.want {
			t.Errorf("longestSubarray(%v) = %d; want %d", tt.nums, result, tt.want)
		}
	}
}
