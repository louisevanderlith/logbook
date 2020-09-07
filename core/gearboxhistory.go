package core

import (
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type GearboxHistory struct {
	SwapDate time.Time
	SerialNo string
	//Gearbox  *Gearbox
}

func (m GearboxHistory) Valid() error {
	return validation.Struct(m)
}
