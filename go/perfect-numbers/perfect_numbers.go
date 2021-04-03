package perfect

import (
	"errors"
)

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

type Classification int

var ErrOnlyPositive = errors.New("only possitive values")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	} else if n == 1 {
		return ClassificationDeficient, nil
	}

	var sum, div int64 = 1, 2
	for div*div < n {
		if n%div == 0 {
			sum += div
			if n/div != n {
				sum += n / div
			}
		}
		div++
	}

	if sum > n {
		return ClassificationAbundant, nil
	} else if sum < n {
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
