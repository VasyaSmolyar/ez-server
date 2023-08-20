package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"ex-server/internal/auth/response"
	"ex-server/internal/task/adaptor"
	"ex-server/internal/task/agent"
)

func HandleError(err error, w http.ResponseWriter) {
	if errors.Is(err, adaptor.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else if errors.Is(err, agent.ErrFileNotFound) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&response.ErrorResponse{
			Msg: fmt.Sprintln(err),
		})
	} else if errors.Is(err, agent.ErrForbidden) {
		w.WriteHeader(http.StatusForbidden)
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
