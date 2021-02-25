package rotationalcipher

// RotationalCipher return encrypted input
func RotationalCipher(input string, shift int) string {
	result := make([]byte, len(input))
	byteShift := byte(shift)

	for i := 0; i < len(input); i++ {
		if input[i] >= 'a' && input[i] <= 'z' {
			if input[i]+byteShift > 'z' {
				result[i] = input[i] + byteShift - 26
			} else {
				result[i] = input[i] + byteShift
			}
		} else if input[i] >= 'A' && input[i] <= 'Z' {
			if input[i]+byteShift > 'Z' {
				result[i] = input[i] + byteShift - 26
			} else {
				result[i] = input[i] + byteShift
			}
		} else {
			result[i] = input[i]
		}

	}

	return string(result)
}
