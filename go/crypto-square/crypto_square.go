package cryptosquare

import "strings"

// Encode encrypt text and return cipher
func Encode(text string) string {
	text = strings.ToLower(text)
	var normalizedBuilder strings.Builder

	if text == "" {
		return text
	}

	for _, mark := range text {
		if (mark >= 'a' && mark <= 'z') || (mark >= '0' && mark <= '9') {
			normalizedBuilder.WriteRune(mark)
		}
	}
	normalized := normalizedBuilder.String()

	textLenght, c, r := len(normalized), len(normalized), 1

	for (c < r) || (c-r > 1) {
		r++
		c = textLenght / r

		if textLenght%r > 0 {
			c++
		}
	}

	result := make([]byte, c*r+c-1)
	var curC, curR int

	for i := 0; i < c*r; i++ {
		index := curC*r + curR
		if index > 0 {
			if index%r == 0 {
				result[index+index/r-1] = ' '
			}
			index = index + index/r
		}

		if i >= len(normalized) {
			result[index] = ' '
		} else {
			result[index] = (normalized[i])
		}

		if curC == c-1 {
			curC = 0
			curR++
		} else {
			curC++
		}
	}

	return string(result)
}
