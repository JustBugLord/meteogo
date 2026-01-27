package langs

func GetWMCDescription(wmc int, lang Language) (string, bool) {
	switch lang {
	case English:
		return GetEnglishWMC(wmc)
	case Russian:
		return GetRussianWMC(wmc)
	}
	return "", false
}

func GetWindDirectionByLang(direction int, lang Language) (string, bool) {
	windDirection := GetWindDirection(direction)
	switch lang {
	case English:
		return GetEnglishNameDirection(windDirection)
	case Russian:
		return GetRussianNameDirection(windDirection)
	}
	return "", false
}
