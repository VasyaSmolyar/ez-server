package handler

import (
	"errors"
	"log"
	"net/http"

	"ex-server/internal/task/adaptor"
)

func HandleError(err error, w http.ResponseWriter) {
	if errors.Is(err, adaptor.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
