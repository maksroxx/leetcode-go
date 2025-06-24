package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	first := make(map[byte]string)
	second := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		if word, exist := first[pattern[i]]; exist {
			if word != words[i] {
				return false
			}
		} else {
			first[pattern[i]] = words[i]
		}
		if char, exist := second[words[i]]; exist {
			if char != pattern[i] {
				return false
			}
		} else {
			second[words[i]] = pattern[i]
		}
	}
	return true
}

func wordPatternSol(pattern string, s string) bool {

	words := strings.Fields(s)
	sLen := len(words)

	if sLen != len(pattern) {
		return false
	}

	Map1 := make(map[string]string)
	Map2 := make(map[string]string)

	for i, word := range words {
		p := string(pattern[i])
		w := string(word)
		if Map1[p] != "" && Map1[p] != w {
			return false
		}

		if Map2[w] != "" && Map2[w] != p {
			return false
		}
		Map1[p] = w
		Map2[w] = p
	}

	return true
}
