package groupanagrams

import (
	"fmt"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)
	for _, str := range strs {
		count := make([]int, 26)
		for _, c := range str {
			count[c-'a']++
		}
		key := makeKey(count)
		anagramGroups[key] = append(anagramGroups[key], str)
	}
	result := [][]string{}
	for _, group := range anagramGroups {
		result = append(result, group)
	}
	return result
}

func makeKey(count []int) string {
	var sb strings.Builder
	for _, c := range count {
		sb.WriteString(fmt.Sprintf("%d#", c))
	}
	return sb.String()
}

// O(n * k)
// n str count
// k len str
