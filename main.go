package main

import (
	"log"
	"os"
	"path"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/logbook/core"
	"github.com/louisevanderlith/logbook/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	appName := beego.BConfig.AppName

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	srv := mango.NewService(appName, pubPath, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv, host)

		beego.Run()
	}
}
