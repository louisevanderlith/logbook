package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/logbook/core"
)

// /v1/history/:vehicleKey
func ViewHistory(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	vehKey := ctx.FindParam("key")

	key, err := husk.ParseKey(vehKey)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.GetHistoryByVehicle(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// /v1/history
func CreateHistory(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	var obj core.History
	err := ctx.Body(&obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
