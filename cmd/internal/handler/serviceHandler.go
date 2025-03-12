package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Prettyletto/service-notifier/cmd/internal/model"
	"github.com/Prettyletto/service-notifier/cmd/internal/service"
)

type ServiceHandler struct {
	service service.ServiceService
}

func NewServiceHandler(service service.ServiceService) *ServiceHandler {
	return &ServiceHandler{service: service}
}

func (h *ServiceHandler) CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var service model.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Error in the payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateService(&service); err != nil {
		http.Error(w, fmt.Sprintf("Faield to create service: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ServiceHandler) GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	services, err := h.service.RetrieveAllServices()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retireve services: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(services); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ServiceHandler) GetServiceByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	service, err := h.service.RetrieveServiceById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve servicee :%v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(service); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *ServiceHandler) UpdateServiceHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var newService model.Service
	if err := json.NewDecoder(r.Body).Decode(&newService); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode payload: %v", err), http.StatusBadRequest)
	}

	updated, err := h.service.UpdateService(id, &newService)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update service: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	updated.ID = id

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode service: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *ServiceHandler) DeleteCollectionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.service.DeleteService(id); err != nil {
		http.Error(w, fmt.Sprintf("failed to delete collection: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}
