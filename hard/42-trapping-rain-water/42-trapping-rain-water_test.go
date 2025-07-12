package trappingrainwater

import "testing"

func TestTrappingRainWater(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{}, 0},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, tt := range tests {
		result := trap(tt.height)
		if result != tt.want {
			t.Errorf("trap(%v) = %d; want %d", tt.height, result, tt.want)
		}
	}
}
