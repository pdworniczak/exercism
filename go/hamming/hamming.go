package hamming

import (
	"errors"
)

/*
Distance function compate two DNA sequences, and calculate "hamming distance".
Function take two string value representing DNA sequences and return number of different characters that occured between this two strings.
*/
func Distance(a, b string) (int, error) {
	ra := []rune(a)
	rb := []rune(b)

	runeCount := len(ra)

	if runeCount != len(rb) {
		return 0, errors.New("dna sequences have different length")
	}

	distance := 0

	for i := 0; i < runeCount; i++ {
		if ra[i] != rb[i] {
			distance++
		}
	}

	return distance, nil
}
