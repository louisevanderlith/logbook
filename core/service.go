package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

type Service struct {
	UserKey    hsk.Key
	Odometer   int64
	LicensePlate string
	OilChange  bool
	Inspection bool
	Additional Additional
}

func (o Service) Valid() error {
	return validation.Struct(o)
}
