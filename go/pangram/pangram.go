package pangram

import (
	"strings"
)

// IsPangram check if input string use all letters
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for letter := 'a'; letter <= 'z'; letter++ {
		if !strings.ContainsRune(s, letter) {
			return false
		}
	}
	return true
}
