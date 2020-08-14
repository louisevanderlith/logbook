package core

import (
	"github.com/louisevanderlith/husk"
)

//History holds the entire history of a vehicle.
type History struct {
	VehicleKey husk.Key
	ClientKey  husk.Key
	Services   []Service
}

func (p History) Valid() error {
	return husk.ValidateStruct(p)
}

//GetHistoryByVehicle finds the history for a given vehicle key.
func GetHistoryByVehicle(vehicleKey husk.Key) (husk.Recorder, error) {
	return ctx.History.FindFirst(byVehicleKey(vehicleKey))
}

func (p History) Create() (husk.Recorder, error) {
	rec, err := ctx.History.Create(p)

	if err != nil {
		return nil, err
	}

	err = ctx.History.Save()

	if err != nil {
		return nil, err
	}

	return rec, nil
}
