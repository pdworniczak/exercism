package wordcount

import (
	"strings"
	"unicode"
)

// Frequency result
type Frequency map[string]int

// WordCount counts words in input
func WordCount(input string) (result Frequency) {
	result = make(Frequency)
	rinput := []rune(strings.ToLower(input))

	stop, start := 0, 0
	for nstop, r := range rinput {
		stop = nstop
		if isWhitemark(r) {
			for ; start < len(rinput) && isWhitemark(rinput[start]); start++ {
			}

			if start < stop {
				result[getKey(rinput, start, stop)]++
				start = stop + 1
			}
		}
	}

	if len(rinput[start:]) > 0 {
		result[getLastKey(rinput, start)]++
	}

	return
}

func isWhitemark(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '\''
}

func getKey(input []rune, start, stop int) string {
	return strings.Trim(string(input[start:stop]), "'")
}

func getLastKey(input []rune, start int) string {
	return strings.Trim(string(input[start:]), "'")
}
