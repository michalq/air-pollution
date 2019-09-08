package airly

import (
	"github.com/go-openapi/strfmt"
	"github.com/michalq/go-airly-api-client/client"
)

type Configuration struct {
	Host     string
	BasePath string
}

func NewClient(configuration Configuration) *client.AirlyAPIClient {
	transport := &client.TransportConfig{
		Host:     configuration.Host,
		BasePath: configuration.BasePath,
		Schemes:  []string{"http"},
	}

	return client.NewHTTPClientWithConfig(strfmt.NewFormats(), transport)
}
