package app

import (
	"crud-example/config"
	"crud-example/internal/db"
	"crud-example/internal/tpl"
	"fmt"
	"net/http"
)

type cfg struct{}

func Run() error {
	app := cfg{}
	// Initialize env variables
	config.LoadEnv()
	// Initialize templates
	tpl.LoadTemplates()
	// Initialize mongodb
	db.InitMongoClient()
	// Close MongoDB connection on exit
	defer db.DisconnectMongoClient()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.WebPort),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
