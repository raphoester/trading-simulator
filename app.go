package main

import (
	"net/http"

	"trading-simulator/internal/simulator"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) Run() {
	app.Router = mux.NewRouter()
	app.RouteInit()
	http.ListenAndServe(":8080", app.Router)
	fmt.Println("Server started on port 8080")
}

func (app *App) RouteInit() {
	global := app.Router.PathPrefix("/api/v1").Subrouter()
	global.HandleFunc("/simulate", simulator.SimulatorHandler).Methods("GET")
}
