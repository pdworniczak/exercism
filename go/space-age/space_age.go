package space

type Planet string

const (
	Mercury = "Mercury"
	Venus = "Venus"
	Earth = "Earth"
	Mars = "Mars"
	Jupiter = "Jupiter"
	Saturn = "Saturn"
	Uranus = "Uranus"
	Neptune = "Neptune"
)
	

/*
Mercury: orbital period 0.2408467 Earth years
Venus: orbital period 0.61519726 Earth years
Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
Mars: orbital period 1.8808158 Earth years
Jupiter: orbital period 11.862615 Earth years
Saturn: orbital period 29.447498 Earth years
Uranus: orbital period 84.016846 Earth years
Neptune: orbital period 164.79132 Earth years
*/	

func Age(seconds float64, planet Planet) float64  {
	yeasrOnEarth := seconds / 31557600

	switch planet {	
	case Mercury: 
		return yeasrOnEarth * 1 / 0.2408467
	case Venus:
		return yeasrOnEarth * 1 /  0.61519726
	case Earth:
		return yeasrOnEarth
	case Mars:
		return yeasrOnEarth * 1 /  1.8808158
	case Jupiter:
		return yeasrOnEarth * 1 /  11.862615
	case Saturn:
		return yeasrOnEarth * 1 /  29.447498
	case Uranus:
		return yeasrOnEarth * 1 /  84.016846
	case Neptune:
		return yeasrOnEarth * 1 /  164.79132
	}
	return 0
}
