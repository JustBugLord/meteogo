package langs

type Language string

const (
	English Language = "en"
	Russian Language = "ru"
)

type WindDirection string

const (
	Unknown   WindDirection = ""
	North     WindDirection = "n"
	NorthEast WindDirection = "ne"
	East      WindDirection = "e"
	SouthEast WindDirection = "se"
	South     WindDirection = "s"
	SouthWest WindDirection = "sw"
	West      WindDirection = "w"
	NorthWest WindDirection = "nw"
)

func GetWindDirection(grads int) WindDirection {
	if grads < 0 {
		grads *= -1
	}
	if (grads >= 335 && grads <= 360) || (grads >= 0 && grads <= 25) {
		return North
	} else if grads >= 25 && grads <= 65 {
		return NorthEast
	} else if grads >= 65 && grads <= 115 {
		return East
	} else if grads >= 115 && grads <= 155 {
		return SouthEast
	} else if grads >= 155 && grads <= 205 {
		return South
	} else if grads >= 205 && grads <= 245 {
		return SouthWest
	} else if grads >= 245 && grads <= 295 {
		return West
	} else if grads >= 295 && grads <= 335 {
		return NorthWest
	}
	return Unknown
}
