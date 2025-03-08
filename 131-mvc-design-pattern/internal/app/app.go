package app

import (
	"mvc-design-pattern/internal/controller"
	"mvc-design-pattern/internal/view"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Run() error {
	// Initialize templates
	view.LoadTemplates()

	router := httprouter.New()
	uc := controller.NewUserController()
	router.GET("/", uc.GetHome)
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)
	return http.ListenAndServe(":8080", router)
}
