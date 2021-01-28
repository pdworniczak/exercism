package luhn

import (
	"strings"
)

// Valid check if luhn code is valid
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	inputLen := len(input)

	if inputLen < 2 {
		return false
	}

	var sum, i, val int

	for _, mark := range input {
		if mark >= '0' && mark <= '9' {
			val = int(mark) - '0'
			if (inputLen%2 != 0 && i%2 != 0) || (inputLen%2 == 0 && i%2 == 0) {
				val = val * 2
				if val > 9 {
					val -= 9
				}
			}
			sum += val
			i++
		} else {
			return false
		}
	}

	return sum%10 == 0
}
