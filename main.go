package main

import (
	"air-pollution/config"
	"air-pollution/modules/core"
	coreRepositories "air-pollution/modules/core/repositories"
	"air-pollution/modules/providers/airly"
	"air-pollution/modules/providers/system"
	"log"

	"air-pollution/modules/providers/gios"
	"fmt"
)

func main() {
	configuration := config.Build()
	db := core.CreatePostgresConnection(configuration.Postgres)
	acquisitionRepository := system.NewAcquisitionRepository(db)

	giosClient := gios.New(configuration.ProviderGios)
	airlyClient, airlyAuth := airly.New(configuration.ProviderAirly)

	var daqSupervisors []coreRepositories.FinderManager
	daqSupervisors = append(daqSupervisors, coreRepositories.FinderManager{
		IsEnabled:         true,
		StationFinder:     gios.NewStationRepository(giosClient),
		AcquisitionFinder: gios.NewAcquisitionRepository(giosClient),
	})
	daqSupervisors = append(daqSupervisors, coreRepositories.FinderManager{
		IsEnabled:         false,
		StationFinder:     airly.NewStationRepository(airlyClient, airlyAuth),
		AcquisitionFinder: gios.NewAcquisitionRepository(giosClient),
	})

	for _, daqSupervisor := range daqSupervisors {
		if !daqSupervisor.IsEnabled {
			continue
		}
		stations, err := daqSupervisor.StationFinder.FindAll()
		if err != nil {
			log.Fatal("Error", err)
		}
		for _, station := range stations {
			fmt.Println(station.Name, station.ExternalID)
			acquisitions, err := daqSupervisor.AcquisitionFinder.FindAllByStationID(station.ExternalID)
			if err != nil {
				log.Fatal(err)
			}
			for _, acq := range acquisitions {
				err = acquisitionRepository.Save(acq)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	//acqs, err := airlyAcqRepository.FindAllByStationID("8077")
	//fmt.Printf("Fetching acquisitions for stationId 8077, len:  %d", len(acqs))
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//for _, acq := range acqs {
	//	fmt.Printf("Type: %s, Value: %s, Day: %s\n", acq.Type, acq.Value, acq.DateFrom.String())
	//}
	//
	//giosAcqRepository := gios.NewAcquisitionRepository(giosClient)
	//acqs, err = giosAcqRepository.FindAllByStationID("145")
	//fmt.Printf("Fetching acquisitions for stationId 145, len:  %d", len(acqs))
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//for _, acq := range acqs {
	//	fmt.Printf("Type: %s, Value: %s, Day: %d\n", acq.Type, acq.Value, acq.DateFrom.Day())
	//}
}
