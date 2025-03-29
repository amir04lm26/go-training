package app

import (
	"crud-example/api/handlers"
	"crud-example/config"
	"crud-example/internal/db"
	approuter "crud-example/internal/router"
	"crud-example/internal/tpl"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Run() error {
	// Initialize env variables
	config.LoadEnv()
	// Initialize templates
	tpl.LoadTemplates()
	// Initialize mongodb
	db.InitMongoClient()
	// Close MongoDB connection on exit
	defer db.DisconnectMongoClient()

	router := httprouter.New()
	bc := handlers.NewBookController()

	// Set up the routes
	approuter.SetupRoutes(router, bc)

	return http.ListenAndServe(":8080", router)
}
