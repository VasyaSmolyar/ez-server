package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"ex-server/internal/task/adaptor"
	"ex-server/internal/task/handler"
	"ex-server/pkg/config"
	"ex-server/pkg/db"
)

const (
	waitTimeout  = time.Second * 15
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

type Server struct {
	config  *viper.Viper
	Handler *handler.Handler
}

func Init(configPath string) (*Server, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	db, err := db.Init(cfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	repo := adaptor.Init(db.Connection)
	return &Server{config: cfg, Handler: handler.Init(*repo)}, nil
}

func (s *Server) Run() {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.config.GetString("App.Host"), s.config.GetString("App.Port")),
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

	r.HandleFunc("/task/list", s.Handler.GetTasksList).Methods("GET")
	r.HandleFunc("/task/{id}", s.Handler.GetTask).Methods("GET")
	r.HandleFunc("/task", s.Handler.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id}", s.Handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id}", s.Handler.DeleteTask).Methods("DELETE")

	return r
}
