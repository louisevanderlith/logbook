package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type historyFilter func(obj History) bool

func (f historyFilter) Filter(obj hsk.Record) bool {
	return f(obj.Data().(History))
}

func byVehicleKey(key hsk.Key) historyFilter {
	return func(obj History) bool {
		return obj.VehicleKey == key
	}
}
