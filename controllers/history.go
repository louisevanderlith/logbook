package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/logbook/core"
)

type HistoryController struct {
}

// /v1/history/:vehicleKey
func (req *HistoryController) GetByVehicle(ctx context.Contexer) (int, interface{}) {
	vehKey := ctx.FindParam("vehicleKey")

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
func (req *HistoryController) Post(ctx context.Contexer) (int, interface{}) {
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
