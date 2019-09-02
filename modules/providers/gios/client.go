package gios

import (
	"github.com/go-openapi/strfmt"
	"github.com/michalq/go-gios-api-client/client"
)

type Configuration struct {
	Host string
	BasePath string
}

func NewClient(configuration Configuration) *client.GiosAPIClient {
	transport := &client.TransportConfig{
		Host:     configuration.Host,
		BasePath: configuration.BasePath,
		Schemes:  []string{"http"},
	}

	return client.NewHTTPClientWithConfig(strfmt.NewFormats(), transport)
}