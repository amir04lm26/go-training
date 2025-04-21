package app

import (
	"crud-example/api/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *cfg) routes() http.Handler {
	router := httprouter.New()
	bc := handlers.NewBookController()

	router.Handler(http.MethodGet, "/", http.RedirectHandler("/books", http.StatusSeeOther))
	router.GET("/books", bc.GetBooks)
	router.GET("/book/details/:isbn", bc.GetBookDetails)
	router.GET("/book/create", bc.GetCreateBook)
	router.POST("/book/create", bc.CreateBookProcess)
	router.GET("/book/update/:isbn", bc.GetUpdateBook)
	router.PUT("/book/update/:isbn", bc.PutUpdatedBook)
	router.DELETE("/book/delete/:isbn", bc.DeleteBookProcess)

	return router
}
