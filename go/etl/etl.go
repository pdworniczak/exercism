package etl

import "strings"

// Transform transforam input data to output
func Transform(data map[int][]string) map[string]int {
	result := make(map[string]int)

	for value, keys := range data {
		for _, key := range keys {
			result[strings.ToLower(key)] = value
		}
	}

	return result
}
