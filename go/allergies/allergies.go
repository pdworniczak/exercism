package allergies

type res struct {
	substance string
	result    bool
}

var substances = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score uint) []string {
	var result = []string{}

	for index, allergie := range substances {
		if 1<<index&score == 1<<index {
			result = append(result, allergie)
		}
	}

	return result
}

func AllergicTo(score uint, substance string) bool {
	for index, allergie := range substances {
		if substance == allergie && 1<<index&score == 1<<index {
			return true
		}
	}
	return false
}
