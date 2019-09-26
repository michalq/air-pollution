package repositories

type FinderManager struct {
	IsEnabled         bool
	StationFinder     StationFinder
	AcquisitionFinder AcquisitionFinder
}
