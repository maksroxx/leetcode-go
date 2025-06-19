package validpalindrom

import (
	"reflect"
	"testing"
)

func TestValidPalindrom(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"aba", true},
		{"", true},
		{"&adda&", true},
		{"aboba", true},
		{"Maks", false},
	}
	for _, tt := range tests {
		result := isPalindrome(tt.str)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("isPalindrome(%s) = %t; want %t", tt.str, result, tt.want)
		}
	}
}
