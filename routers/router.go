package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/logbook/controllers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/mango/enums"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	beego.Router("/", controllers.NewDefaultCtrl(ctrlmap))
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)
	emptyMap["GET"] = enums.User

	ctrlmap.Add("/", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
