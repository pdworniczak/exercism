package cipher

import (
	"strings"
)

type cesar struct {
}

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

const alphabetCount = ('z' - 'a' + 1)

func (s *vigenere) Encode(text string) string {
	var result strings.Builder
	keyLen, i := len(s.key), 0
	text = strings.ToLower(text)

	for _, mark := range text {
		if mark >= 'a' && mark <= 'z' {
			result.WriteRune((mark+rune(s.key[i%keyLen]-'a')-'a'+alphabetCount)%alphabetCount + 'a')
			i++
		}
	}

	return result.String()
}

func (s *vigenere) Decode(text string) string {
	var result strings.Builder
	keyLen := len(s.key)

	for i, mark := range text {
		if mark >= 'a' && mark <= 'z' {
			result.WriteRune((mark-rune(s.key[i%keyLen])+alphabetCount)%alphabetCount + 'a')
		}
	}

	return result.String()
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance >= 26 || distance <= -26 || distance == 0 {
		return nil
	}
	return NewVigenere(string('a' + (distance+26)%26))
}

func NewVigenere(key string) Cipher {
	if keyIsValid(key) {
		return &vigenere{key}
	}

	return nil
}

func keyIsValid(key string) bool {
	if key == "" {
		return false
	}

	isValid := false

	for _, mark := range key {
		if mark > 'z' || mark < 'a' {
			return false
		}

		if mark != 'a' {
			isValid = true
		}
	}

	return isValid
}
