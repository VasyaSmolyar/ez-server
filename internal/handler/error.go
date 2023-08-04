package handler

import (
	"errors"
	"ex-server/internal/adaptor"
	"log"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter) {
	if errors.Is(err, adaptor.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
