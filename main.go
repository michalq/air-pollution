package main

import (
	"air-pollution/config"
	coreRepositories "air-pollution/modules/core/repositories"
	"air-pollution/modules/providers/airly"
	airlyRepositories "air-pollution/modules/providers/airly/repositories"
	"github.com/jpfuentes2/go-env"
	"log"
	"os"
	"path"
	"strings"

	"air-pollution/modules/providers/gios"
	giosRepositories "air-pollution/modules/providers/gios/repositories"
	"fmt"
)

func main() {

	configuration := &config.Config{}
	pwd, err := os.Getwd()
	env.ReadEnv(path.Join(pwd, ".env"))
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range os.Environ() {
		line := strings.Split(v, "=")
		switch line[0] {
		case "AIRLY_API_CLIENT":
			configuration.Airly.AuthKey = strings.Join(line[1:], "")
		}
	}

	var stationsRepositories []coreRepositories.StationRepositoryInterface
	giosClient := gios.New(gios.Configuration{
		Host:     "api.gios.gov.pl",
		BasePath: "/pjp-api/rest",
	})
	stationsRepositories = append(stationsRepositories, giosRepositories.NewStationRepository(giosClient))

	airlyClient, airlyAuth := airly.New(airly.Configuration{
		Host:     "airapi.airly.eu",
		BasePath: "/",
		AuthKey:  configuration.Airly.AuthKey,
	})
	stationsRepositories = append(stationsRepositories, airlyRepositories.NewStationRepository(airlyClient, airlyAuth))

	for _, stationRepository := range stationsRepositories {
		stations, err := stationRepository.FindAll()
		if err != nil {
			log.Fatal("Error", err)
		}
		for _, station := range stations {
			fmt.Println(station.Name)
		}
	}
}
