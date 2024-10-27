package main

import (
	"coffee-app/adapters"
	"coffee-app/core"
)

func main() {
	heater := adapters.NewElectricHeater()
	pump := adapters.NewThermosiphon(heater)
	maker := core.NewCoffeeMaker(pump, heater)
	maker.Brew()
}
