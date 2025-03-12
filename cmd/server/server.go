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
		log.Fatal(initErr)
	}
	companyRepo := repository.NewCompanyRepo(database)
	companyService := service.NewCompanyService(companyRepo)
	companyHandler := handler.NewCompanyHandler(companyService)

	clientRepo := repository.NewClientRepo(database, companyRepo)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(clientService)

	serviceRepo := repository.NewServiceRepo(database)
	serviceService := service.NewServiceService(serviceRepo)
	serviceHandler := handler.NewServiceHandler(serviceService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /companies/{id}/clients", clientHandler.CreateClientHandler)
	mux.HandleFunc("GET /clients", clientHandler.GetAllClientsHandler)
	mux.HandleFunc("GET /clients/{id}", clientHandler.GetClientByIdHandler)
	mux.HandleFunc("PUT /clients/{id}", clientHandler.UpdateClientHandler)
	mux.HandleFunc("DELETE /clients/{id}", clientHandler.DeleteCollectionHandler)

	mux.HandleFunc("POST /companies", companyHandler.CreateCompanyHandler)
	mux.HandleFunc("GET /companies", companyHandler.GetAllCompaniesHandler)
	mux.HandleFunc("GET /companies/{id}", companyHandler.GetCompanyByIdHandler)
	mux.HandleFunc("PUT /companies/{id}", companyHandler.UpdateCompanyHandler)
	mux.HandleFunc("DELETE /companies/{id}", companyHandler.DeleteCollectionHandler)

	mux.HandleFunc("POST /services", serviceHandler.CreateServiceHandler)
	mux.HandleFunc("GET /services", serviceHandler.GetAllServicesHandler)
	mux.HandleFunc("GET /services/{id}", serviceHandler.GetServiceByIdHandler)
	mux.HandleFunc("PUT /services/{id}", serviceHandler.UpdateServiceHandler)
	mux.HandleFunc("DELETE /services/{id}", serviceHandler.DeleteCollectionHandler)

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
