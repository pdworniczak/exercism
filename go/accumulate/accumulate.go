package accumulate

type converter func(string) string

// Accumulate take given string array and call given converter function on all elements of array and return string array with converter results
func Accumulate(words []string, cv converter) []string {
	result := make([]string, len(words))

	for i, word := range words {
		result[i] = cv(word)
	}

	return result
}
