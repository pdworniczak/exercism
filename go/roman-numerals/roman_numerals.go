package romannumerals

import (
	"fmt"
	"strings"
)

var (
	roman = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
)

// ToRomanNumeral convert arabic numbert roman
func ToRomanNumeral(num int) (string, error) {
	if num < 1 || num > 3000 {
		return "", fmt.Errorf("forbiden input value %d", num)
	}

	result := ""

	for _, v := range []int{1, 10, 100, 1000} {
		a := num % 10
		switch a {
		case 1, 2, 3:
			result = strings.Repeat(roman[v], a) + result
		case 4:
			result = roman[v] + roman[v*5] + result
		case 5:
			result = roman[v*5] + result
		case 6, 7, 8:
			result = roman[v*5] + strings.Repeat(roman[v], a-5) + result
		case 9:
			result = roman[v] + roman[v*10] + result
		}
		num /= 10
	}

	return result, nil
}
