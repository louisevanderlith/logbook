package main

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/logbook/core"
	"github.com/louisevanderlith/logbook/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
)

func main() {
	mode := os.Getenv("RUNMODE")
	appName := beego.BConfig.AppName

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	srv := mango.NewService(mode, appName, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)

		beego.Run()
	}
}
