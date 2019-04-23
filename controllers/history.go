package controllers

import (
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
		req.Serve(nil, err)
		return
	}

	req.Serve(core.GetHistoryByVehicle(vehKey))
}

func (req *HistoryController) Post() {
	
}