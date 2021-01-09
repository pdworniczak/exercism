package lsproduct

import (
	"errors"
)

// LargestSeriesProduct return highest sum of san digits sereis in digits
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span == 0 {
		return 1, nil
	} else if span < 1 {
		return -1, errors.New("span must be positive")
	}

	if span > len(digits) {
		return 0, errors.New("span is longer than digits we have")
	}

	listOfDigits, convertError := convertDigits(digits)

	if convertError != nil {
		return -1, convertError
	}

	result := 0

	for i := 0; i+span-1 < len(listOfDigits); i++ {
		var partialResult int
		for j := 0; j < span; j++ {
			if j == 0 {
				partialResult = listOfDigits[j+i]
			} else {
				partialResult *= listOfDigits[j+i]
			}
		}

		if partialResult > result {
			result = partialResult
		}
	}

	return result, nil
}

func convertDigits(digits string) ([]int, error) {
	listOfDigits := make([]int, len(digits))

	for _, r := range []rune(digits) {
		if r >= '0' && r <= '9' {
			listOfDigits = append(listOfDigits, int(r-'0'))
		} else {
			return nil, errors.New("Can't convert digits")
		}
	}

	return listOfDigits, nil
}
