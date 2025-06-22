package canconstruct

import (
	"reflect"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		r    string
		m    string
		want bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aa", true},
	}
	for _, tt := range tests {
		result := canConstruct(tt.r, tt.m)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("canConstruct(%s, %s) = %t; want %t", tt.r, tt.m, result, tt.want)
		}
	}
}
