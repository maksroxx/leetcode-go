package issubsequence

func isSubsequence(s string, t string) bool {
	first, second := 0, 0
	for second < len(t) && first < len(s) {
		if s[first] == t[second] {
			first++
		}
		second++
	}
	return first == len(s)
}
