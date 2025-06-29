package stringcompression

import "testing"

func TestCompress(t *testing.T) {
	tests := []struct {
		chars []byte
		want  int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{[]byte{'a'}, 1},
		{[]byte{'a', 'b'}, 2},
		{[]byte{'a', 'a', 'a', 'a'}, 2},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b'}, 3},
		{[]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}, 6},
	}
	for _, tt := range tests {
		t.Run(string(tt.chars), func(t *testing.T) {
			got := compress(tt.chars)
			if got != tt.want {
				t.Errorf("compress(%+v) = %d, want %d", tt.chars, got, tt.want)
			}
		})
	}
}
