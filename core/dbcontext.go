package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	History husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		History: husk.NewTable(History{}),
	}
}

func Shutdown() {
	ctx.History.Save()
}
