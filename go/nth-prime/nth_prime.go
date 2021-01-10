package prime

import (
	"math"
)

// Nth return nth prime number
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	prime := 2
	primesCount := 0

	for i := prime; primesCount < n; {
		if isPrime(i) {
			primesCount++
			prime = i
		}
		i++
	}

	return prime, true
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
