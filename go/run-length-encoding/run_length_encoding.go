package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode encode string
func RunLengthEncode(input string) string {
	runeInput := []rune(input)
	var mark rune
	countMarks := 1

	var result strings.Builder

	if len(runeInput) == 0 {
		return ""
	}

	mark = runeInput[0]

	for i := 1; i < len(runeInput); i++ {
		if mark != runeInput[i] {
			if countMarks != 1 {
				result.Write([]byte(strconv.Itoa(countMarks)))
				countMarks = 1
			}
			result.WriteRune(mark)
			mark = runeInput[i]
		} else {
			countMarks++
		}
	}

	if countMarks > 1 {
		result.WriteRune(rune(int('0' + countMarks)))
	}

	result.WriteRune(mark)

	return result.String()
}

// RunLengthDecode decode string
func RunLengthDecode(input string) string {
	count := int64(0)
	var result strings.Builder

	for _, r := range input {
		if unicode.IsNumber(r) {
			if count > 0 {
				if v, err := strconv.ParseInt(string(r), 10, 0); err == nil {
					count = count*10 + v
				}
			} else {
				count, _ = strconv.ParseInt(string(r), 10, 0)
			}
		} else {
			if count == 0 {
				result.WriteRune(r)
			} else {
				result.WriteString(strings.Repeat(string(r), int(count)))
				count = 0
			}
		}
	}

	return result.String()
}
