package maxprofit

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{}, 0},
	}
	for _, tt := range tests {
		result := maxProfit(tt.prices)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("maxProfit(%v) = %d; want %d", tt.prices, result, tt.expected)
		}
	}
}
