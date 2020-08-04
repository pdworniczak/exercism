package anagram

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

//Detect check if it is possible to create candite word from given input
func Detect(input string, candidates []string) []string {
	result := make([]string, 0)

	letters := strings.Split(strings.ToLower(input), "")

	sort.Strings(letters)

	fmt.Println(input, letters)
	for _, candidate := range candidates {
		if strings.ToLower(input) == strings.ToLower(candidate) {
			continue
		}

		candidateLetters := strings.Split(strings.ToLower(candidate), "")
		sort.Strings(candidateLetters)
		if reflect.DeepEqual(letters, candidateLetters) {
			result = append(result, candidate)
		}
	}

	return result
}
