package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	History husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		History: husk.NewTable(History{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.History.Save()
}
