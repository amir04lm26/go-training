package apperr

import (
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// Define custom error variables
var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrInternalServer = errors.New("internal server error")
)
var ErrNoMessage = ""

func HandleHttpMongoErr(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		http.NotFound(w, r)
	case errors.Is(err, ErrNoItemFoundToDelete):
		http.NotFound(w, r)
	case errors.Is(err, ErrNoItemFoundToUpdate):
		http.NotFound(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func HandleBadRequest(w http.ResponseWriter, message string) {
	http.Error(w, http.StatusText(http.StatusBadRequest)+" "+message, http.StatusBadRequest)
}

func HandleNotAcceptable(w http.ResponseWriter, message string) {
	http.Error(w, http.StatusText(http.StatusNotAcceptable)+" "+message, http.StatusNotAcceptable)
}

func HandleInternalServerError(w http.ResponseWriter, message string) {
	http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+message, http.StatusInternalServerError)
}
