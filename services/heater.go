package services

type Heater interface {
	Off()
	On()
	IsHot() bool
}
