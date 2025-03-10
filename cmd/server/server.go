package server

import (
	"log"
	"net/http"

	"github.com/Prettyletto/service-notifier/cmd/internal/db"
	"github.com/Prettyletto/service-notifier/cmd/internal/handler"
	"github.com/Prettyletto/service-notifier/cmd/internal/repository"
	"github.com/Prettyletto/service-notifier/cmd/internal/service"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string) *Server {
	database, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	initErr := database.Init()
	if initErr != nil {
		log.Fatal(err)
	}

	clientRepo := repository.NewClientRepo(database)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(clientService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /clients", clientHandler.CreateClientHandler)
	mux.HandleFunc("GET /clients", clientHandler.GetAllClientsHandler)
	mux.HandleFunc("GET /clients/{id}", clientHandler.GetClientByIdHandler)
	mux.HandleFunc("PUT /clients/{id}", clientHandler.UpdateClientHandler)
	mux.HandleFunc("DELETE /clients/{id}", clientHandler.DeleteCollectionHandler)

	s := &Server{httpServer: &http.Server{
		Addr:    addr,
		Handler: mux,
	}}

	return s
}

func (s *Server) Start() {
	log.Println("Server listening in port:", s.httpServer.Addr)
	s.httpServer.ListenAndServe()
}
func (s *Server) Stop() {
	log.Println("Closing the server...")
	s.httpServer.Close()
}
