package main

import (
	coreRepositories "air-pollution/modules/core/repositories"
	"air-pollution/modules/providers/airly"
	airlyRepositories "air-pollution/modules/providers/airly/repositories"
	//"air-pollution/modules/providers/gios"
	//giosRepositories "air-pollution/modules/providers/gios/repositories"
	"fmt"
)

func main() {

	var stationsRepositories []coreRepositories.StationRepositoryInterface
	//giosClient := gios.NewClient(gios.Configuration{
	//	Host:     "api.gios.gov.pl",
	//	BasePath: "/pjp-api/rest",
	//})
	//stationsRepositories = append(stationsRepositories, giosRepositories.NewStationRepository(giosClient))

	airlyClient := airly.NewClient(airly.Configuration{
		Host:     "https://airapi.airly.eu",
		BasePath: "/",
	})
	stationsRepositories = append(stationsRepositories, airlyRepositories.NewStationRepository(airlyClient))

	for _, stationRepository := range stationsRepositories {
		stations, _ := stationRepository.FindAll()
		for _, station := range stations {
			fmt.Println(station.Address.City)
		}
	}
}
