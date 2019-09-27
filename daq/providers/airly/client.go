package airly

import (
	"github.com/go-openapi/runtime"
	openApiClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/michalq/go-airly-api-client/client"
)

type Configuration struct {
	Host     string
	BasePath string
	AuthKey  string
}

func New(configuration Configuration) (*client.AirlyAPIClient, runtime.ClientAuthInfoWriter) {
	transport := &client.TransportConfig{
		Host:     configuration.Host,
		BasePath: configuration.BasePath,
		Schemes:  []string{"http"},
	}
	auth := openApiClient.APIKeyAuth("apikey", "header", configuration.AuthKey)
	return client.NewHTTPClientWithConfig(strfmt.NewFormats(), transport), auth
}
