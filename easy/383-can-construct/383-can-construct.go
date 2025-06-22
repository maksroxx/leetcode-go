package canconstruct

func canConstruct(ransomNote string, magazine string) bool {
	mapa := make(map[rune]int)
	for _, char := range magazine {
		mapa[char]++
	}
	for _, char := range ransomNote {
		if mapa[char] == 0 {
			return false
		}
		mapa[char]--
	}
	return true
}
