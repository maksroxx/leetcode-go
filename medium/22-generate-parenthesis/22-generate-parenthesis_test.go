package generateparenthesis

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{
			n:        1,
			expected: []string{"()"},
		},
		{
			n:        2,
			expected: []string{"(())", "()()"},
		},
		{
			n: 3,
			expected: []string{
				"((()))", "(()())", "(())()", "()(())", "()()()",
			},
		},
		{
			n:        0,
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		result := generateParenthesis(tt.n)

		sort.Strings(result)
		sort.Strings(tt.expected)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("generateParenthesis(%d) = %v; expected %v", tt.n, result, tt.expected)
		}
	}
}
