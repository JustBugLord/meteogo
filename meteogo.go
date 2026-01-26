package meteogo

import "github.com/JustBugLord/reqtango/v2"

type Meteo struct {
	rb *reqtango.RequestBuilder
}

func NewMeteo() *Meteo {
	return &Meteo{
		rb: reqtango.NewRequestBuilderSimple(),
	}
}
