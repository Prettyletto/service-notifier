package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Prettyletto/service-notifier/cmd/internal/model"
	"github.com/Prettyletto/service-notifier/cmd/internal/service"
)

type ClientHandler struct {
	service service.ClientService
}

func NewClientHandler(service service.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h *ClientHandler) CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	var client model.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Error in the payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateClient(&client); err != nil {
		http.Error(w, fmt.Sprintf("Faield to create client: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ClientHandler) GetAllClientsHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := h.service.RetrieveAllClients()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retireve clients: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(clients); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ClientHandler) GetClientByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	client, err := h.service.RetrieveClientById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve cliente :%v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *ClientHandler) UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var newClient model.Client
	if err := json.NewDecoder(r.Body).Decode(&newClient); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode payload: %v", err), http.StatusBadRequest)
	}

	updated, err := h.service.UpdateClient(id, &newClient)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update client: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	updated.ID = id

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode client: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *ClientHandler) DeleteCollectionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.service.DeleteClient(id); err != nil {
		http.Error(w, fmt.Sprintf("failed to delete collection: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}
