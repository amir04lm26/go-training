package app

import (
	"mongo-basic/internal/controller"
	"mongo-basic/internal/db"
	"mongo-basic/internal/view"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Run() error {
	// Initialize templates
	view.LoadTemplates()
	client := db.GetMongoClient()
	// Close MongoDB connection on exit
	defer db.DisconnectMongoClient(client)
	router := httprouter.New()
	uc := controller.NewUserController(client)

	router.GET("/", uc.GetHome)
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	return http.ListenAndServe(":8080", router)
}
