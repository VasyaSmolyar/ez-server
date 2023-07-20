package main

import (
	"ex-server/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/Tasks", handlers.GetTasksList).Methods("GET")
	r.HandleFunc("/Tasks/{id}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/Tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/Tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/Tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
