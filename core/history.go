package core

import (
	"github.com/louisevanderlith/husk"
)

type History struct {
	VehicleKey husk.Key
	Contact    ContactDetail
	Services   []Service
}
