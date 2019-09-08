package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/logbook/core"
)

type History struct {
}

func (x *History) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

func (x *History) Search(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

// /v1/history/:vehicleKey
func (req *History) View(ctx context.Requester) (int, interface{}) {
	vehKey := ctx.FindParam("key")

	key, err := husk.ParseKey(vehKey)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := core.GetHistoryByVehicle(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, rec
}

// /v1/history
func (req *History) Create(ctx context.Requester) (int, interface{}) {
	var obj core.History
	err := ctx.Body(&obj)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := obj.Create()

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, rec
}
