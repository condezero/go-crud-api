package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()
	r.Get("/hc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "healthy")
	})
	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

func (serv *Server) Close() error {
	// TODO: add resource closure.
	return nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
