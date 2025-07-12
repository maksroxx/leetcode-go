package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			intervals: [][]int{},
			want:      [][]int{},
		},
		{
			intervals: [][]int{{1, 4}, {0, 4}},
			want:      [][]int{{0, 4}},
		},
	}
	for _, tt := range tests {
		result := merge(tt.intervals)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("merge(%v) = %v; want %v", tt.intervals, result, tt.want)
		}
	}
}
