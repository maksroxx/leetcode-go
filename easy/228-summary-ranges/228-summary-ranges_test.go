package summaryranges

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		input []int
		want  []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{}, []string{}},
		{[]int{5}, []string{"5"}},
		{[]int{1, 3}, []string{"1", "3"}},
	}

	for _, tt := range tests {
		actual := summaryRanges(tt.input)
		if !reflect.DeepEqual(actual, tt.want) {
			t.Errorf("summaryRanges(%v) = %v; expected %v", tt.input, actual, tt.want)
		}
	}
}
