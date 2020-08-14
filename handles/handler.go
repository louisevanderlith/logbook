package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()

	view := kong.ResourceMiddleware(http.DefaultClient, "logbook.history.view", scrt, securityUrl, managerUrl, ViewHistory)
	r.HandleFunc("/history/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	create := kong.ResourceMiddleware(http.DefaultClient, "logbook.history.create", scrt, securityUrl, managerUrl, CreateHistory)
	r.HandleFunc("/history", create).Methods(http.MethodPost)

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "logbook.history.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
