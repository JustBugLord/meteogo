package hourly

type Units struct {
	Time                     string `json:"time"`
	Temperature2M            string `json:"temperature_2m"`
	RelativeHumidity2M       string `json:"relative_humidity_2m"`
	WeatherCode              string `json:"weather_code"`
	ApparentTemperature      string `json:"apparent_temperature"`
	Evapotranspiration       string `json:"evapotranspiration"`
	Visibility               string `json:"visibility"`
	DewPoint2M               string `json:"dew_point_2m"`
	Precipitation            string `json:"precipitation"`
	PrecipitationProbability string `json:"precipitation_probability"`
}
type Data struct {
	Time                     []string  `json:"time"`
	Temperature2M            []float64 `json:"temperature_2m"`
	RelativeHumidity2M       []int     `json:"relative_humidity_2m"`
	WeatherCode              []int     `json:"weather_code"`
	ApparentTemperature      []float64 `json:"apparent_temperature"`
	Evapotranspiration       []float64 `json:"evapotranspiration"`
	Visibility               []float64 `json:"visibility"`
	DewPoint2M               []float64 `json:"dew_point_2m"`
	Precipitation            []float64 `json:"precipitation"`
	PrecipitationProbability []int     `json:"precipitation_probability"`
}
