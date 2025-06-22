package issubsequence

import (
	"reflect"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, tt := range tests {
		result := isSubsequence(tt.s, tt.t)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("isSubsequence(%s, %s) = %t; want %t", tt.s, tt.t, result, tt.want)
		}
	}
}
