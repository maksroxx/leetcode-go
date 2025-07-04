package checkinclusion

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"ab", "a", false},
		{"a", "ab", true},
	}
	for _, tt := range tests {
		result := checkInclusion(tt.s1, tt.s2)
		if result != tt.want {
			t.Errorf("checkInclusion(%s, %s) = %t; want %t", tt.s1, tt.s2, result, tt.want)
		}
	}
}
