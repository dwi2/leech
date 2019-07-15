package validpalindrome

import "strings"

func inRange(r byte) bool {
	return (r >= 48 && r <= 57) ||
		(r >= 97 && r <= 122) ||
		(r >= 65 && r <= 90)
}

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	lowS := strings.ToLower(s)
	result := true
	head := 0
	tail := len(s) - 1

	for head < tail {
		if !inRange(lowS[head]) {
			head++
		} else if !inRange(lowS[tail]) {
			tail--
		} else if lowS[head] != lowS[tail] {
			result = false
			return result
		} else {
			head++
			tail--
		}
	}

	return result
}
