package daily

type Units struct {
	Time             string `json:"time"`
	WeatherCode      string `json:"weather_code"`
	Temperature2MMax string `json:"temperature_2m_max"`
	Temperature2MMin string `json:"temperature_2m_min"`
}
type Data struct {
	Time             []string  `json:"time"`
	WeatherCode      []int     `json:"weather_code"`
	Temperature2MMax []float64 `json:"temperature_2m_max"`
	Temperature2MMin []float64 `json:"temperature_2m_min"`
}
