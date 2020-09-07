package core

import (
	"github.com/louisevanderlith/husk/validation"
	"time"

	"github.com/louisevanderlith/logbook/core/alterationtype"
)

type Alteration struct {
	AlterDate   time.Time
	AlterType   alterationtype.Enum
	Code        string
	Description string
	Odometer    int64
}

func (m Alteration) Valid() error {
	return validation.Struct(m)
}
