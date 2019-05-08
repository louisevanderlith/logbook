package core

import (
	"github.com/louisevanderlith/husk"
)

//History holds the entire history of a vehicle.
type History struct {
	VehicleKey husk.Key
	Contact    ContactDetail
	Services   []Service
}

func (p History) Valid() (bool, error) {
	return husk.ValidateStruct(&p)
}

//GetHistoryByVehicle finds the history for a given vehicle key.
func GetHistoryByVehicle(vehicleKey husk.Key) (husk.Recorder, error) {
	return ctx.History.FindFirst(byVehicleKey(vehicleKey))
}

func (p History) Create() (husk.Recorder, error) {
	rec := ctx.History.Create(p)

	if rec.Error != nil {
		return nil, rec.Error
	}

	ctx.History.Save()

	return rec.Record, nil
}
