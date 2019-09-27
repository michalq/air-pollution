package repositories

import (
	"air-pollution/daq/core/models"
)

type FinderSupervisor struct {
	IsEnabled         bool
	StationFinder     StationFinder
	AcquisitionFinder AcquisitionFinder
}

type FinderSupervisors struct {
	stationFinders     []StationFinder
	acquisitionFinders []AcquisitionFinder
}

func NewFinderSupervisors() *FinderSupervisors {
	return &FinderSupervisors{}
}

func (f *FinderSupervisors) Add(supervisors ...FinderSupervisor) {
	for _, supervisor := range supervisors {
		if !supervisor.IsEnabled {
			continue
		}
		f.stationFinders = append(f.stationFinders, supervisor.StationFinder)
		f.acquisitionFinders = append(f.acquisitionFinders, supervisor.AcquisitionFinder)
	}
}

func (f *FinderSupervisors) FindAllStations() ([]*models.Station, error) {
	stations := make([]*models.Station, 0)
	var tmpStations []*models.Station
	var err error
	for _, stationFinder := range f.stationFinders {
		tmpStations, err = stationFinder.FindAll()
		if err != nil {
			return stations, err
		}
		stations = append(stations, tmpStations...)
	}
	return stations, nil
}

func (f *FinderSupervisors) FindAllAcquisitionsByStationID(stationID string) ([]*models.Acquisition, error) {
	return nil, nil
}
