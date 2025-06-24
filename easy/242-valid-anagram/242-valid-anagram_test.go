package validanagram

import (
	"reflect"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"aba", "ab", false},
		{"rat", "car", false},
	}
	for _, tt := range tests {
		result := isAnagram(tt.s, tt.t)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("isAnagram(%s, %s) = %t; want %t", tt.s, tt.t, result, tt.want)
		}
	}
}
