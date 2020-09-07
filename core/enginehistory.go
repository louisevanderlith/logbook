package core

import (
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type EngineHistory struct {
	SwapDate time.Time
	SerialNo string
	//Engine   *Engine
}

func (m EngineHistory) Valid() error {
	return validation.Struct(m)
}
