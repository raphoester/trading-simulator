package main

import (
	"net/http"

	"trading-simulator/internal/simulator"
	"fmt"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) Run() {
	app.Router = mux.NewRouter()
	app.RouteInit()
	fmt.Println("Started server on port 8080")
	if err := http.ListenAndServe(":8080", app.Router); err != nil {
		fmt.Printf("Failed starting server | %s", err.Error())
	} 
}

func (app *App) RouteInit() {
	global := app.Router.PathPrefix("/api/v1").Subrouter()
	global.HandleFunc("/simulate", simulator.SimulatorHandler).Methods("GET")
}
