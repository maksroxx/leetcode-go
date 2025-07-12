package maxdisttoclosest

import "testing"

func TestMaxDistToClosest(t *testing.T) {
	tests := []struct {
		seats []int
		want  int
	}{
		{[]int{1, 0, 0, 0, 1, 0, 1}, 2},
		{[]int{1, 0, 0, 0}, 3},
		{[]int{0, 1}, 1},
	}
	for _, tt := range tests {
		result := maxDistToClosest(tt.seats)
		if result != tt.want {
			t.Errorf("maxDistToClosest(%v) = %d; want %d", tt.seats, result, tt.want)
		}
	}
}
