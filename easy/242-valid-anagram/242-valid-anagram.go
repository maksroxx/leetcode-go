package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sArr := [26]int{}
	tArr := [26]int{}
	for _, tt := range s {
		sArr[tt-'a']++
	}
	for _, tt := range t {
		tArr[tt-'a']++
	}
	return sArr == tArr
}
