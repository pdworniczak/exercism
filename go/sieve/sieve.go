package sieve

func Sieve(limit int) []int {
	marked := make([]bool, limit-1)
	result := make([]int, 0)

	sieveLength := len(marked)
	for i := 0; i < sieveLength; i++ {
		if marked[i] == false {
			result = append(result, i+2)
			for j := 2*i + 2; j < sieveLength; j += i + 2 {
				marked[j] = true
			}
		}
	}

	return result
}
