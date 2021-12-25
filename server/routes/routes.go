package routes

import (
	"tripit/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.GetHome).Methods("GET", "OPTIONS")

	//TODO...
	//router.HandleFunc("/signup", ...).Methods("POST", "OPTIONS")
	//router.HandleFunc("/signin", ...).Methods("POST", "OPTIONS")
	//router.HandleFunc("/logout", ...).Methods("POST", "OPTIONS")
	//router.HandleFunc("/user/{id}", ...).Methods("GET", "OPTIONS")
	//router.HandleFunc("/user/{id}", ...).Methods("PUT", "OPTIONS")
	//router.HandleFunc("/user/{id}", ...).Methods("DELETE", "OPTIONS")

	return router
}
