package core

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type GearboxHistory struct {
	SwapDate time.Time
	SerialNo string
	//Gearbox  *Gearbox
}

func (m GearboxHistory) Valid() error {
	return husk.ValidateStruct(&m)
}
