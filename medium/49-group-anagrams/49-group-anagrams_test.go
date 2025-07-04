package groupanagrams

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func equalGroupAnagrams(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	normalize := func(groups [][]string) []string {
		normalized := make([]string, len(groups))
		for i, group := range groups {
			sort.Strings(group)
			normalized[i] = strings.Join(group, ",")
		}
		sort.Strings(normalized)
		return normalized
	}

	return reflect.DeepEqual(normalize(a), normalize(b))
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			input:    []string{"abc", "cba", "bca", "xyz", "zyx"},
			expected: [][]string{{"abc", "cba", "bca"}, {"xyz", "zyx"}},
		},
		{
			input:    []string{"listen", "silent", "enlist", "google", "gogole"},
			expected: [][]string{{"listen", "silent", "enlist"}, {"google", "gogole"}},
		},
		{
			input:    []string{},
			expected: [][]string{},
		},
	}

	for _, tt := range tests {
		result := groupAnagrams(tt.input)
		if !equalGroupAnagrams(result, tt.expected) {
			t.Errorf("groupAnagrams(%v) = %v; expected %v", tt.input, result, tt.expected)
		}
	}
}
