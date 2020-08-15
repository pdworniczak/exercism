package wordcount

import (
	"strings"
	"unicode"
)

// Frequency result
type Frequency map[string]int

// WordCount counts words in input
func WordCount(input string) Frequency {
	result := make(Frequency)
	rinput := []rune(strings.ToLower(input))

	start := 0
	for stop, r := range rinput {
		if isWhitemark(r) {
			if r == '\'' && len(rinput) > stop+1 && unicode.IsLetter(rinput[stop+1]) {
				continue
			}
			for ; start < len(rinput) && isWhitemark(rinput[start]); start++ {
			}
			if start < stop {
				result[string(rinput[start:stop])]++
				start = stop + 1
			}
		}
	}

	if len(rinput[start:]) > 0 {
		result[string(rinput[start:])]++
	}

	return result
}

func isWhitemark(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}
