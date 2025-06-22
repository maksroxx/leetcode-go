package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if mapped, exist := sMap[s[i]]; exist {
			if mapped != t[i] {
				return false
			}
		} else {
			sMap[s[i]] = t[i]
		}
		if mapped, exist := tMap[t[i]]; exist {
			if mapped != s[i] {
				return false
			}
		} else {
			tMap[t[i]] = s[i]
		}
	}
	return true
}
