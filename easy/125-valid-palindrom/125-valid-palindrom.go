package validpalindrom

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !isGood(rune(s[left])) {
			left++
		}
		for left < right && !isGood(rune(s[right])) {
			right--
		}
		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isGood(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
