package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && !hasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func hasPrefix(str, prefix string) bool {
	if len(prefix) > len(str) {
		return false
	}
	return str[:len(prefix)] == prefix
}
