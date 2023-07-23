package server

import (
	"context"
	"ex-server/internal/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const (
	waitTimeout  = time.Second * 15
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

type Server struct {
	port    int
	Handler *handler.Handler
}

func Init(port int, handler *handler.Handler) *Server {
	return &Server{port: port, Handler: handler}
}

func (s *Server) Run() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      s.initRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), waitTimeout)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
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
