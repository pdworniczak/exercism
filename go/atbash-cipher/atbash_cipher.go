package atbash

import (
	"strings"
)

// Atbash encrypt input text
func Atbash(text string) string {
	var result strings.Builder
	var i int

	text = strings.ToLower(text)

	for _, mark := range text {
		if i == 5 {
			result.WriteRune(' ')
			i = 0
		}

		if mark >= 'a' && mark <= 'z' {
			result.WriteRune('z' - (mark - 'a'))
		} else if mark >= '0' && mark <= '9' {
			result.WriteRune(mark)
		} else {
			continue
		}
		i++
	}

	return strings.Trim(result.String(), " ")
}
