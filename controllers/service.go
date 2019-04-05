package controllers

import "github.com/louisevanderlith/mango/control"

type ServiceController struct {
	control.UIController
}

func NewServiceController(ctrlMap *control.ControllerMap) *ServiceController {
	result := &ServiceController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ServiceController) Get() {
	c.Setup("service")

}
