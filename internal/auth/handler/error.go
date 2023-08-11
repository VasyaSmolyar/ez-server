package handler

import (
	"encoding/json"
	"errors"
	"ex-server/internal/auth/exception"
	"ex-server/internal/auth/response"
	"fmt"
	"log"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter) {
	if errors.Is(err, exception.ErrWrongCreds) || errors.Is(err, exception.ErrAlreadyCreated) ||
		errors.Is(err, exception.ErrTokenExpired) || errors.Is(err, exception.ErrTokenInvalid) ||
		errors.Is(err, exception.ErrHashingFailed) {
		w.WriteHeader(http.StatusForbidden)
	} else if errors.Is(err, exception.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(&response.ErrorResponse{
		Msg: fmt.Sprintln(err),
	})
}
