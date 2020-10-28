package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	r.Handle("/history/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewHistory))).Methods(http.MethodGet)
	r.Handle("/history", mw.Handler(http.HandlerFunc(CreateHistory))).Methods(http.MethodPost)

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "logbook.history.view", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
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
