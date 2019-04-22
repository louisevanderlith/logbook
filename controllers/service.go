package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type ServiceController struct {
	control.UIController
}

func NewServiceController(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *ServiceController {
	result := &ServiceController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ServiceController) Get() {
	c.Setup("service")

}
