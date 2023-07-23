package server

import (
	"ex-server/internal/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port    int
	Handler *handler.Handler
}

func Init(port int, handler *handler.Handler) *Server {
	return &Server{port: port, Handler: handler}
}

func (s *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.initRouter())
}

func (s *Server) initRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/Tasks", s.Handler.GetTasksList).Methods("GET")
	r.HandleFunc("/Tasks/{id}", s.Handler.GetTask).Methods("GET")
	r.HandleFunc("/Tasks", s.Handler.CreateTask).Methods("POST")
	r.HandleFunc("/Tasks/{id}", s.Handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/Tasks/{id}", s.Handler.DeleteTask).Methods("DELETE")

	return r
}
