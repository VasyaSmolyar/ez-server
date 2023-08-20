package handler

import (
	"errors"
	"ex-server/internal/object/adaptor"
	"log"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter) {
	if errors.Is(err, adaptor.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else if errors.Is(err, ErrBadRequest) {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
