package strand

// ToRNA function transform dna code to rna
func ToRNA(dna string) string {
	result := make([]rune, len(dna))

	for i, char := range dna {
		switch char {
		case 'G':
			result[i] = 'C'
		case 'C':
			result[i] = 'G'
		case 'T':
			result[i] = 'A'
		case 'A':
			result[i] = 'U'
		}
	}

	return string(result)
}
