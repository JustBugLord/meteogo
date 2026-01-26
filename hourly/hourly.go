package hourly

type Hourly string

const (
	WeatherCode         Hourly = "weather_code"
	Temperature2M       Hourly = "temperature_2m"
	RelativeHumidity2M  Hourly = "relative_humidity_2m"
	ApparentTemperature Hourly = "apparent_temperature"
	Evapotranspiration  Hourly = "evapotranspiration"
	Visibility          Hourly = "visibility"
)

func (h Hourly) String() string {
	return string(h)
}
