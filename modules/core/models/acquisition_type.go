package models

type Type string

const (
	TypePM1  Type = "PM1"
	TypePM10 Type = "PM10"
	TypePM25 Type = "PM25"

	TypeNO2  Type = "NO2"
	TypeSO2  Type = "SO2"
	TypeC6H6 Type = "C6H6"
	TypeO3   Type = "O3"
	TypeCO   Type = "CO"

	TypePressure    Type = "Pressure"
	TypeTemperature Type = "Temperature"
)
