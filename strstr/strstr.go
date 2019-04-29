package strstr

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	result := -1
	for i := 0; i < len(haystack) && result < 0; i++ {
		for j := 0; j < len(needle) && result < 0; j++ {
			if needle[j] != haystack[i+j] {
				result = -1
				break
			} else if needle[j] == haystack[i+j] && j == len(needle)-1 {
				result = i
				break
			} else {
				result = i
			}
		}
	}

	return result
}
