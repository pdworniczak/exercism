package isbn

// IsValidISBN check if isbn 10 is valid
func IsValidISBN(isbn string) bool {
	sum := 0
	i := 0

	for _, mark := range isbn {
		if i > 9 {
			return false
		} else if mark >= '0' && mark <= '9' {
			sum += (int(mark) - '0') * (10 - i)
			i++
		} else if mark == 'x' || mark == 'X' {
			if i == 9 {
				sum += 10 * (10 - i)
				i++
			} else {
				return false
			}
		} else if mark == '-' {
			continue
		} else {
			return false
		}

	}

	return i == 10 && sum%11 == 0
}
