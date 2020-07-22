package raindrops

import (
	"strconv"
)

type raindrop struct {
	divider int
	sound   string
}

var (
	raindrops = []raindrop{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}
)

// Convert function get int as a parametr and convert it to rain sounds it gives, returning as a string value
func Convert(num int) string {
	result := ""

	for _, raindrop := range raindrops {
		if num%raindrop.divider == 0 {
			result += raindrop.sound
		}
	}

	if len(result) == 0 {
		return strconv.Itoa(num)
	}

	return result
}
