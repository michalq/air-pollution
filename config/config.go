package config

import (
	"air-pollution/modules/core"
	"air-pollution/modules/providers/airly"
	"air-pollution/modules/providers/gios"
	"github.com/jpfuentes2/go-env"
	"log"
	"os"
	"path"
	"strings"
)

type Config struct {
	ProviderAirly airly.Configuration
	ProviderGios  gios.Configuration
	Postgres      core.PostgresConfig
}

func Build() *Config {
	configuration := &Config{}
	pwd, err := os.Getwd()
	env.ReadEnv(path.Join(pwd, ".env"))
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range os.Environ() {
		line := strings.Split(v, "=")
		value := strings.Join(line[1:], "")
		switch line[0] {
		case "AIRLY_API_KEY":
			configuration.ProviderAirly.AuthKey = value
		case "AIRLY_HOST":
			configuration.ProviderAirly.Host = value
		case "AIRLY_BASE_PATH":
			configuration.ProviderAirly.BasePath = value
		case "GIOS_HOST":
			configuration.ProviderGios.Host = value
		case "GIOS_BASE_PATH":
			configuration.ProviderGios.BasePath = value
		}
	}

	return configuration
}
