package romantoint

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"", 0},
		{"I", 1},
		{"D", 500},
		{"O", 0},
		{"XL", 40},
		{"XC", 90},
		{"MMXXIII", 2023},
	}
	for _, tt := range tests {
		result := romanToInt(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("romanToInt(%q) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}
