package adapters

import "coffee-app/services"

type ElectricHeater struct {
	heating bool
}

func (e *ElectricHeater) Off() {
	e.heating = false
}

func (e *ElectricHeater) On() {
	println("~ ~ ~ ~ heating ~ ~ ~ ~")
	e.heating = true
}

func (e *ElectricHeater) IsHot() bool {
	return e.heating
}

func NewElectricHeater() services.Heater {
	return &ElectricHeater{}
}
