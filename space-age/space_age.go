package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var unit float64 = 31557600
	var val = seconds / unit
	var ratio = 1.0
	switch planet {
	case "Mercury":
		ratio = 0.2408467
	case "Venus":
		ratio = 0.61519726
	case "Earth":
		ratio = 1.0
	case "Mars":
		ratio = 1.8808158
	case "Jupiter":
		ratio = 11.862615
	case "Saturn":
		ratio = 29.447498
	case "Uranus":
		ratio = 84.016846
	case "Neptune":
		ratio = 164.79132
	default:
	}

	return val / ratio
}
