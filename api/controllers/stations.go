package controllers

import "air-pollution/api"

type StationsController struct{}

func NewStationsController() *StationsController {
	return &StationsController{}
}

func (*StationsController) GetStations() (*api.HttpContext, error) {
	return nil, nil
}
