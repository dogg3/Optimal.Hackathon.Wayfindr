package main

// Wrapper interface enforces required functionality for wrapped devices
type Wrapper interface {
	GetID() string
	GetSourceLongitude() float64
	GetSourceLatitude() float64
	GetTargetLongitude() float64
	GetTargetLatitude() float64
	CalcDistance() float64
	CalcSpeed() float64
}
