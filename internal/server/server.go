package server

import (
	"context"
	"ex-server/internal/adaptor"
	"ex-server/internal/entity"
	"ex-server/internal/handler"
	"ex-server/pkg/config"
	"ex-server/pkg/db"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
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

func Init(port int, configPath string) (*Server, error) {
	db, err := initDB(configPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	taskRepo := &adaptor.TaskRepository{}

	handler := handler.Init(db, *taskRepo)

	return &Server{port: port, Handler: handler}, nil
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

	r.HandleFunc("/task/list", s.Handler.GetTasksList).Methods("GET")
	r.HandleFunc("/task/{id}", s.Handler.GetTask).Methods("GET")
	r.HandleFunc("/task", s.Handler.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id}", s.Handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id}", s.Handler.DeleteTask).Methods("DELETE")

	return r
}

func initDB(filename string) (*gorm.DB, error) {
	cfg, err := config.LoadConfig(filename)
	if err != nil {
		return nil, err
	}

	return db.NewConnection(cfg,
		&entity.Task{},
	)
}
