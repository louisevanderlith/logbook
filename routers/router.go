package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/logbook/controllers"
)

func Setup(poxy *droxolite.Epoxy) {
	//History
	histCtrl := &controllers.HistoryController{}
	histGroup := droxolite.NewRouteGroup("history", histCtrl)
	histGroup.AddRoute("Create Service History", "/", "POST", roletype.Owner, histCtrl.Post)
	histGroup.AddRoute("History by Vehicle", "/{vehicleKey:[0-9]+\x60[0-9]+}", "GET", roletype.User, histCtrl.GetByVehicle)
	poxy.AddGroup(histGroup)
	/*ctrlmap := EnableFilter(s, host)

	beego.Router("/v1/history", controllers.NewHistoryCtrl(ctrlmap))*/
}

/*
func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.User
	emptyMap["POST"] = roletype.Owner
	emptyMap["PUT"] = roletype.Owner

	ctrlmap.Add("/v1/history", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}*/
