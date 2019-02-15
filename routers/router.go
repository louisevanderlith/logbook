package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/logbook/controllers"
	"github.com/louisevanderlith/mango/pkg"
	"github.com/louisevanderlith/mango/pkg/control"
	"github.com/louisevanderlith/mango/pkg/enums"
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
