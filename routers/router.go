package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/logbook/controllers"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.Owner, mix.JSON, &controllers.History{})
	//History
	/*
		histCtrl := &controllers.HistoryController{}
		histGroup := routing.NewRouteGroup("history", mix.JSON)
		histGroup.AddRoute("Create Service History", "", "POST", roletype.Owner, histCtrl.Post)
		histGroup.AddRoute("History by Vehicle", "/{vehicleKey:[0-9]+\x60[0-9]+}", "GET", roletype.User, histCtrl.GetByVehicle)
		e.AddBundle(histGroup)
	*/
}
