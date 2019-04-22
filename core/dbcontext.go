package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	History husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		History: husk.NewTable(new(History)),
	}
}

func Shutdown() {
	ctx.History.Save()
}
