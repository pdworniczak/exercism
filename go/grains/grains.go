package grains

import (
	"errors"
)

// Square number of grains on specific square
func Square(squares int) (uint64, error) {
	if squares < 1 || squares > 64 {
		return 0, errors.New("wrong squares input")
	}
	return 1 << (squares - 1), nil
}

// Total number of grains on chesse board
func Total() uint64 {
	return 1<<64 - 1
}
