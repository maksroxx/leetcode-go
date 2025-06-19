package longestcommonprefix

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{}, ""},
	}
	for _, tt := range tests {
		result := longestCommonPrefix(tt.words)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("longestCommonPrefix(%v) = %s; want %s", tt.words, result, tt.want)
		}
	}
}
