package protein

import (
	"errors"
)

var (
	// ErrStop error connected with specific codon
	ErrStop = errors.New("STOP")
	// ErrInvalidBase error for all invalid codon
	ErrInvalidBase = errors.New("Invalid base protei")
	rnaprotein     = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
)

// FromCodon translate rna codon to protein, throws error when get STOP codon or invalid codon parameter
func FromCodon(codon string) (string, error) {
	if protein, ok := rnaprotein[codon]; ok {
		if protein == "STOP" {
			return "", ErrStop
		}
		return protein, nil
	}

	return "", ErrInvalidBase
}

// FromRNA translate rna codos into proteins, when STOP codon occures it imediatly return what was allready tranlated, when invalid base occure return error
func FromRNA(codons string) ([]string, error) {
	result := []string{}
	var reterr error

	codon := ""
	for _, letter := range codons {
		codon += string(letter)
		if len(codon) == 3 {
			protein, err := FromCodon(codon)
			if err != nil {
				if err != ErrStop {
					reterr = err
				}
				break
			}
			result = append(result, protein)
			codon = ""
		}
	}

	return result, reterr
}
