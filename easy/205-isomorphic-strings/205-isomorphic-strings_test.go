package isomorphicstrings

import (
	"reflect"
	"testing"
)

func TestIsomorphicStrings(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"badc", "baba", false},
		{"a", "bc", false},
	}
	for _, tt := range tests {
		result := isIsomorphic(tt.s, tt.t)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("isIsomorphic(%s, %s) = %t; want %t", tt.s, tt.s, result, tt.want)
		}
	}
}
