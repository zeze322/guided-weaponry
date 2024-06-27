package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zeze322/wt-guided-weaponry/storage"
)

type Server struct {
	port  string
	mysql storage.Storage
	mongo storage.Storage
}

func NewServer(port string, mysql, mongo storage.Storage) *Server {
	return &Server{
		port:  port,
		mysql: mysql,
		mongo: mongo,
	}
}

func (s *Server) Start() {
	router := chi.NewMux()

	router.Get("/category", makeHTTP(s.handleListCategory))
	router.Get("/category/{category}", makeHTTP(s.handleWeapons))
	router.Get("/{weapon}", makeHTTP(s.handleWeapon))
	router.Post("/add", makeHTTP(s.handleInsertWeapon))
	router.Put("/update", makeHTTP(s.handleUpdateWeapon))

	log.Printf("Server starting on port %s", s.port)
	log.Fatal(http.ListenAndServe(s.port, router))
}
