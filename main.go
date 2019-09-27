package main

import (
	"air-pollution/app"
	"air-pollution/daq/providers/airly"
	"air-pollution/daq/providers/system"
	coreRepositories "air-pollution/daq/repositories"
	"log"

	"air-pollution/daq/providers/gios"
	"fmt"
)

func main() {
	configuration := app.Build()
	db := app.CreatePostgresConnection(configuration.Postgres)
	acquisitionRepository := system.NewAcquisitionRepository(db)

	giosClient := gios.New(configuration.ProviderGios)
	airlyClient, airlyAuth := airly.New(configuration.ProviderAirly)

	var daqSupervisors = coreRepositories.NewFinderSupervisors()
	daqSupervisors.Add(
		coreRepositories.FinderSupervisor{
			IsEnabled:         true,
			StationFinder:     gios.NewStationRepository(giosClient),
			AcquisitionFinder: gios.NewAcquisitionRepository(giosClient),
		},
		coreRepositories.FinderSupervisor{
			IsEnabled:         false,
			StationFinder:     airly.NewStationRepository(airlyClient, airlyAuth),
			AcquisitionFinder: airly.NewAcquisitionRepository(airlyClient, airlyAuth),
		},
	)

	stations, err := daqSupervisors.FindAllStations()
	if err != nil {
		log.Fatal("Error", err)
	}

	for _, station := range stations {
		fmt.Println(station.Name, station.ExternalID)
		acquisitions, err := daqSupervisors.FindAllAcquisitionsByStationID(station.ExternalID)
		if err != nil {
			log.Fatal(err)
		}
		for _, acq := range acquisitions {
			err = acquisitionRepository.Persist(acq)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
