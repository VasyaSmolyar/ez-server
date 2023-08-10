package handler

import (
	"errors"
	"ex-server/internal/auth/exception"
	"log"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter) {
	if errors.Is(err, exception.ErrWrongCreds) || errors.Is(err, exception.ErrAlreadyCreated) {
		w.WriteHeader(http.StatusForbidden)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
