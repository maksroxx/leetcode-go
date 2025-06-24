package validparentheses

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"(){}[]", true},
		{"({})", true},
		{"[(]]", false},
		{"", true},
		{"123", false},
	}
	for _, tt := range tests {
		result := isValid(tt.s)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("isValid(%s) = %t; want %t", tt.s, result, tt.want)
		}
	}
}
