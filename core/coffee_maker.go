package core

import "coffee-app/services"

type CoffeeMaker struct {
	pump   services.Pump
	heater services.Heater
}

func (maker *CoffeeMaker) Brew() {
	maker.heater.On()
	maker.pump.Pumping()
	println("[_]P coffee! [_]P")
	maker.heater.Off()
}

func NewCoffeeMaker(pump services.Pump, heater services.Heater) *CoffeeMaker {
	return &CoffeeMaker{pump: pump, heater: heater}
}
