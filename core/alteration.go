package core

import (
	"time"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/logbook/core/alterationtype"
)

type Alteration struct {
	AlterDate   time.Time
	AlterType   alterationtype.Enum
	Code        string
	Description string
	Odometer    int64
}

func (m Alteration) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
