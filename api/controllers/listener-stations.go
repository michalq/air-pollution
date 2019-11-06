package controllers

import "air-pollution/api"

type AddStationRequestBody struct{}

type ListenerStationsController struct{}

func NewListenerStationsController() *ListenerStationsController {
	return &ListenerStationsController{}
}

func (*ListenerStationsController) AddStation(stationID int, requestBody AddStationRequestBody) (*api.HttpContext, error) {
	return nil, nil
}
func (*ListenerStationsController) DeleteStation(stationID int) (*api.HttpContext, error) {
	return nil, nil
}

func (*ListenerStationsController) GetStations() (*api.HttpContext, error) {
	return nil, nil
}
