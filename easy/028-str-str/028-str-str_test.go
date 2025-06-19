package strstr

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"abc", "", 0},
	}
	for _, tt := range tests {
		result := strStr(tt.haystack, tt.needle)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("strStr(%s, %s) = %d; want %d", tt.haystack, tt.needle, result, tt.want)
		}
	}
}
