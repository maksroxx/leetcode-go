package lengthlastword

import (
	"reflect"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		word string
		want int
	}{
		{"Hello World", 5},
		{"   ", 0},
		{"Go is awesome", 7},
		{"", 0},
	}
	for _, tt := range tests {
		result := lengthOfLastWord(tt.word)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("length of last word(%s) = %d; want %d", tt.word, result, tt.want)
		}
	}
}
