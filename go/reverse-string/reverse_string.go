package reverse

// Reverse take input sstring and revers all runes inside
func Reverse(input string) string {
	inputRunes := []rune(input)
	inputLength := len(inputRunes)
	result := make([]rune, inputLength)

	for i, char := range inputRunes {
		result[inputLength-i-1] = char
	}

	return string(result)
}
