package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/logbook/core"
)

type HistoryController struct {
	xontrols.APICtrl
}

// /v1/history/:vehicleKey
func (req *HistoryController) GetByVehicle() {
	vehKey := req.FindParam("vehicleKey")

	key, err := husk.ParseKey(vehKey)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := core.GetHistoryByVehicle(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// /v1/history
func (req *HistoryController) Post() {
	var obj core.History
	err := req.Body(&obj)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
	}

	req.Serve(http.StatusOK, nil, rec)
}
