package core

import (
	"github.com/louisevanderlith/husk"
)

type Service struct {
	UserKey    husk.Key
	Odometer   int64
	LicensePlate string
	OilChange  bool
	Inspection bool
	Additional Additional
}

func (o Service) Valid() error {
	return husk.ValidateStruct(o)
}
