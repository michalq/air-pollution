package main

import (
	"air-pollution/config"
	coreRepositories "air-pollution/modules/core/repositories"
	"air-pollution/modules/providers/airly"
	airlyRepositories "air-pollution/modules/providers/airly/repositories"
	"log"

	"air-pollution/modules/providers/gios"
	giosRepositories "air-pollution/modules/providers/gios/repositories"
	"fmt"
)

func main() {
	configuration := config.Build()
	var stationsRepositories []coreRepositories.StationRepositoryInterface
	giosClient := gios.New(gios.Configuration{
		Host:     configuration.ProviderGios.Host,
		BasePath: configuration.ProviderGios.BasePath,
	})
	stationsRepositories = append(stationsRepositories, giosRepositories.NewStationRepository(giosClient))

	airlyClient, airlyAuth := airly.New(airly.Configuration{
		Host:     configuration.ProviderAirly.Host,
		BasePath: configuration.ProviderAirly.BasePath,
		AuthKey:  configuration.ProviderAirly.AuthKey,
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
