package diffsquares

// Difference give diff
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SquareOfSum calcualtes square of sum
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	sum = sum * sum
	return sum
}

// SumOfSquares to sum of squares
func SumOfSquares(n int) int {
	sumOfSq := 0
	for i := 1; i <= n; i++ {
		sumOfSq += i * i
	}
	return sumOfSq
}
