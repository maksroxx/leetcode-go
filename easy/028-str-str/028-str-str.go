package strstr

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	hLen := len(haystack)
	nLen := len(needle)

	for i := 0; i <= hLen-nLen; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}
