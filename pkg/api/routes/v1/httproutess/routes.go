package httproutess

import (
	"github.com/gorilla/mux"
	"github.com/iitheo/theohubuc/pkg/api/controllers/v1/userscontroller"
	"github.com/iitheo/theohubuc/pkg/api/middleware"
	"github.com/urfave/negroni"
)

func Router() *negroni.Negroni {
	route := mux.NewRouter()

	n := negroni.Classic()
	n.Use(middleware.Cors())
	n.UseHandler(route)

	//BASE ROUTE
	//route.HandleFunc("/v1", homeHandler)

	//*****************
	// FILMS ROUTES
	//*****************
	filmsRoute := route.PathPrefix("/v1/users").Subrouter()
	filmsRoute.HandleFunc("/create", userscontroller.CreateUser).Methods("POST")

	return n
}
