package router

import (
	"github.com/gorilla/mux"
	"github.com/satanCoffee/arithmetic-api/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/add", middleware.Add).Methods("GET", "OPTIONS")
	router.HandleFunc("/sub", middleware.Sub).Methods("GET", "OPTIONS")
	router.HandleFunc("/mul", middleware.Mul).Methods("GET", "OPTIONS")
	router.HandleFunc("/div", middleware.Div).Methods("GET", "OPTIONS")
	return router
}
