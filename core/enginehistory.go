package core

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type EngineHistory struct {
	SwapDate time.Time
	SerialNo string
	//Engine   *Engine
}

func (m EngineHistory) Valid() error {
	return husk.ValidateStruct(&m)
}
