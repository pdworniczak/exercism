package summultiples

// SumMultiples return sum of all numbers that are multiples of list of divisors
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, div := range divisors {
			if div == 0 {
				continue
			}

			if i%div == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
