package current

type Current string

const (
	WeatherCode         Current = "weather_code"
	Temperature2M       Current = "temperature_2m"
	RelativeHumidity2M  Current = "relative_humidity_2m"
	ApparentTemperature Current = "apparent_temperature"
	WindGusts10M        Current = "wind_gusts_10m"
	WindSpeed10M        Current = "wind_speed_10m"
	WindDirection10M    Current = "wind_direction_10m"
	Snowfall            Current = "snowfall"
	Showers             Current = "showers"
	Rain                Current = "rain"
	Precipitation       Current = "precipitation"
	CloudCover          Current = "cloud_cover"
	PressureMSL         Current = "pressure_msl"
	SurfacePressure     Current = "surface_pressure"
	IsDay               Current = "is_day"
)

func AllCurrents() []Current {
	return []Current{
		WeatherCode,
		Temperature2M,
		RelativeHumidity2M,
		ApparentTemperature,
		WindGusts10M,
		WindSpeed10M,
		WindDirection10M,
		Snowfall,
		Showers,
		Rain,
		Precipitation,
		CloudCover,
		PressureMSL,
		SurfacePressure,
		IsDay,
	}
}

func (c Current) String() string {
	return string(c)
}
