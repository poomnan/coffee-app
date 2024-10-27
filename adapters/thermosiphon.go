package adapters

import "coffee-app/services"

type Thermosiphon struct {
	heater services.Heater
}

func (t *Thermosiphon) Pumping() {
	if t.heater.IsHot() {
		println("=> => => Pumping => => =>")
	}
}

func NewThermosiphon(heater services.Heater) services.Pump {
	return &Thermosiphon{heater: heater}
}
