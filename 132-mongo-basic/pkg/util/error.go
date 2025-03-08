package util

import (
	"mongo-basic/pkg/util/constant"
	"net/http"
)

func InternalServerErr(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", constant.HttpTextPlainContentType)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func BadRequestErr(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", constant.HttpTextPlainContentType)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
