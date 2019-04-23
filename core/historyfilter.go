package core

import "github.com/louisevanderlith/husk"

type historyFilter func(obj *History) bool

func (f historyFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*History))
}

func byVehicleKey(key husk.Key) historyFilter {
	return func(obj *History) bool {
		return obj.VehicleKey == key
	}
}
