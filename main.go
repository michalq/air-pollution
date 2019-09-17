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

type StationRepositorySupervisor struct {
	IsEnabled  bool
	Repository coreRepositories.StationRepositoryInterface
}

func main() {
	configuration := config.Build()
	giosClient := gios.New(gios.Configuration{
		Host:     configuration.ProviderGios.Host,
		BasePath: configuration.ProviderGios.BasePath,
	})

	airlyClient, airlyAuth := airly.New(airly.Configuration{
		Host:     configuration.ProviderAirly.Host,
		BasePath: configuration.ProviderAirly.BasePath,
		AuthKey:  configuration.ProviderAirly.AuthKey,
	})

	var stationsRepositories []StationRepositorySupervisor
	stationsRepositories = append(stationsRepositories, StationRepositorySupervisor{
		IsEnabled:  true,
		Repository: giosRepositories.NewStationRepository(giosClient),
	})
	stationsRepositories = append(stationsRepositories, StationRepositorySupervisor{
		IsEnabled:  false,
		Repository: airlyRepositories.NewStationRepository(airlyClient, airlyAuth),
	})

	for _, stationRepository := range stationsRepositories {
		if !stationRepository.IsEnabled {
			continue
		}
		stations, err := stationRepository.Repository.FindAll()
		if err != nil {
			log.Fatal("Error", err)
		}
		for _, station := range stations {
			fmt.Println(station.Name, station.ExternalID)
		}
	}
}
