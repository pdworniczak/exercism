package pythagorean

import (
	"math"
)

type d struct {
	sum      int
	min, max int
	triplet  Triplet
}

type Triplet struct {
	a int
	b int
	c int
}

// Sum function return all pythagorean triplet that sum of sides is eaqual s paramer
func Sum(s int) []Triplet {
	res := make([]Triplet, 0)
	a := 1
	b := 2
	c := s - (a + b)

	for {
		if pow2(a)+pow2(b) == pow2(c) {
			res = append(res, Triplet{a, b, c})
		}

		b++
		c--

		if b < c {
			continue
		} else {
			a++
			b = a + 1
			c = s - (a + b)

			if a >= b || b >= c {
				break
			}
		}

	}

	return res
}

// Range function return all pythagorean triplet witch sides are in range of min/max parameters
func Range(min, max int) []Triplet {
	res := make([]Triplet, 0)

	for c := max; c >= min+2; c-- {
		for b := c - 1; b >= min+1; b-- {
			for a := b - 1; a >= min; a-- {
				if pow2(a)+pow2(b) == pow2(c) {
					res = append([]Triplet{{a, b, c}}, res...)
				}
			}
		}
	}

	return res
}

func pow2(x int) int {
	return int(math.Pow(float64(x), 2))
}
