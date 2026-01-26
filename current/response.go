package current

type Units struct {
	Time                string `json:"time"`
	Interval            string `json:"interval"`
	Temperature2M       string `json:"temperature_2m"`
	WeatherCode         string `json:"weather_code"`
	WindSpeed10M        string `json:"wind_speed_10m"`
	RelativeHumidity2M  string `json:"relative_humidity_2m"`
	WindDirection10M    string `json:"wind_direction_10m"`
	Snowfall            string `json:"snowfall"`
	Showers             string `json:"showers"`
	Rain                string `json:"rain"`
	Precipitation       string `json:"precipitation"`
	CloudCover          string `json:"cloud_cover"`
	PressureMsl         string `json:"pressure_msl"`
	SurfacePressure     string `json:"surface_pressure"`
	WindGusts10M        string `json:"wind_gusts_10m"`
	ApparentTemperature string `json:"apparent_temperature"`
	IsDay               string `json:"is_day"`
}

type Data struct {
	Time                string  `json:"time"`
	Interval            int     `json:"interval"`
	Temperature2M       float64 `json:"temperature_2m"`
	WeatherCode         int     `json:"weather_code"`
	WindSpeed10M        float64 `json:"wind_speed_10m"`
	RelativeHumidity2M  int     `json:"relative_humidity_2m"`
	WindDirection10M    int     `json:"wind_direction_10m"`
	Snowfall            float64 `json:"snowfall"`
	Showers             float64 `json:"showers"`
	Rain                float64 `json:"rain"`
	Precipitation       float64 `json:"precipitation"`
	CloudCover          int     `json:"cloud_cover"`
	PressureMsl         float64 `json:"pressure_msl"`
	SurfacePressure     float64 `json:"surface_pressure"`
	WindGusts10M        float64 `json:"wind_gusts_10m"`
	ApparentTemperature float64 `json:"apparent_temperature"`
	IsDay               int     `json:"is_day"`
}
