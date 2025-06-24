package happynumber

import (
	"reflect"
	"testing"
)

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		number int
		want   bool
	}{
		{19, true},
		{2, false},
		{1, true},
		{7, false},
	}
	for _, tt := range tests {
		result := isHappy(tt.number)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("isHappy(%d) = %t; want %t", tt.number, result, tt.want)
		}
	}
}
