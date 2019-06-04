package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/logbook/core"
	"github.com/louisevanderlith/mango/control"
)

type HistoryController struct {
	control.APIController
}

func NewHistoryCtrl(ctrlmap *control.ControllerMap) *HistoryController {
	result := &HistoryController{}
	result.SetInstanceMap(ctrlmap)

	return result
}

// /v1/history/:vehicleKey
func (req *HistoryController) GetByVehicle() {
	vehKey := req.Ctx.Input.Param(":vehicleKey")

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
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

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
