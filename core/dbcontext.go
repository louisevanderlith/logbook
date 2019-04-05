package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Vehicles     husk.Tabler
	VINS         husk.Tabler
	Services     husk.Tabler
	ServiceItems husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Vehicles:     husk.NewTable(new(Vehicle)),
		VINS:         husk.NewTable(new(VIN)),
		Services:     husk.NewTable(new(Service)),
		ServiceItems: husk.NewTable(new(ServiceItem)),
	}
}

func Shutdown() {
	ctx.Vehicles.Save()
	ctx.VINS.Save()
	ctx.Services.Save()
	ctx.ServiceItems.Save()
}

func seed() {

}
