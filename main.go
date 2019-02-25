package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/logbook/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
)

func main() {
	mode := beego.BConfig.RunMode
	appName := beego.BConfig.AppName

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
