package diffsquares

// SquareOfSum calculate pow or first n natural numbers
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2

	return sum * sum
}

// SumOfSquares calculate sum of first n natural numbers pow
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference from square of sum substract sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
