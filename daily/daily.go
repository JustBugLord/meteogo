package daily

type Daily string

const (
	WeatherCode      Daily = "weather_code"
	TemperatureMin2M Daily = "temperature_2m_min"
	TemperatureMax2M Daily = "temperature_2m_max"
)

func (d Daily) String() string {
	return string(d)
}
