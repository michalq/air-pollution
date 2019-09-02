package main

import (
	coreRepositories "air-pollution/modules/core/repositories"
	"air-pollution/modules/providers/gios"
	"air-pollution/modules/providers/gios/repositories"
	"fmt"
)

func main() {
	client := gios.NewClient(gios.Configuration{
		Host:     "api.gios.gov.pl",
		BasePath: "/pjp-api/rest",
	})

	var stationsRepositories []coreRepositories.StationRepositoryInterface
	stationsRepositories = append(stationsRepositories, repositories.NewStationRepository(client))

	for _, stationRepository := range stationsRepositories {
		stations, _ := stationRepository.FindAll()
		for _, station := range stations {
			fmt.Println(station.Address.City)
		}
	}
}
