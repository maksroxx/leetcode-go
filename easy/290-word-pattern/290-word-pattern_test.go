package wordpattern

import (
	"reflect"
	"testing"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		first  string
		second string
		want   bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abda", "dog cat cat dog", false},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"a", "dog cat", false},
	}
	for _, tt := range tests {
		result := wordPattern(tt.first, tt.second)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("wordPattern(%s, %s) = %t; want %t", tt.first, tt.second, result, tt.want)
		}
	}
}
