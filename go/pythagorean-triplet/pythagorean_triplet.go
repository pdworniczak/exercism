package pythagorean

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
		if a*a+b*b == c*c {
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

			if b >= c {
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
				if a*a+b*b == c*c {
					res = append([]Triplet{{a, b, c}}, res...)
				}
			}
		}
	}

	return res
}
