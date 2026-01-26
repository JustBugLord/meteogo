package wmc

import (
	ru "meteogo/wmc/en"
	en "meteogo/wmc/ru"
)

func GetWMCDescription(wmc int, lang Language) (string, bool) {
	switch lang {
	case English:
		return en.GetEnglishWMC(wmc)
	case Russian:
		return ru.GetRussianWMC(wmc)
	}
	return "", false
}
