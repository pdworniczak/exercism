package anagram

import (
	"sort"
	"strings"
)

//Detect check if it is possible to create candite word from given input
func Detect(input string, candidates []string) []string {
	result := make([]string, 0)

	inputlc := strings.ToLower(input)
	letters := normalize(inputlc)

	for _, candidate := range candidates {
		candidatelc := strings.ToLower(candidate)
		if inputlc == candidatelc {
			continue
		}

		if letters == normalize(candidatelc) {
			result = append(result, candidate)
		}
	}

	return result
}

func normalize(input string) string {
	result := strings.Split(input, "")
	sort.Strings(result)
	return strings.Join(result, "")
}
