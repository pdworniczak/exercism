package isogram

import "strings"

// IsIsogram test if string parameter is isogram and return bool value
func IsIsogram(sentence string) bool {
	sentence = strings.ToLower(sentence)

	for i, letter := range sentence {
		if !(letter == ' ' || letter == '-') {
			if findLetter(sentence[i+1:], letter) {
				return false
			}
		}
	}
	return true
}

func findLetter(word string, letter rune) bool {
	for _, l := range word {
		if l == letter {
			return true
		}
	}

	return false
}