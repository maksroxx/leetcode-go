package checkinclusion

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count := [26]int{}
	s2Count := [26]int{}

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	if s1Count == s2Count {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		s2Count[s2[i]-'a']++
		s2Count[s2[i-len(s1)]-'a']--
		if s1Count == s2Count {
			return true
		}
	}
	return false
}

// s1 "ab"
// s2 "eidbaoo"
