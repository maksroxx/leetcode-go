package lengthlastword

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return 0
	}
	return len(words[len(words)-1])
}
