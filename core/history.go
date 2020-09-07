package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

//History holds the entire history of a vehicle.
type History struct {
	VehicleKey hsk.Key
	ClientKey  hsk.Key
	Services   []Service
}

func (p History) Valid() error {
	return validation.Struct(p)
}

//GetHistoryByVehicle finds the history for a given vehicle key.
func GetHistoryByVehicle(vehicleKey hsk.Key) (hsk.Record, error) {
	return ctx.History.FindFirst(byVehicleKey(vehicleKey))
}

func (p History) Create() (hsk.Key, error) {
	return ctx.History.Create(p)
}
