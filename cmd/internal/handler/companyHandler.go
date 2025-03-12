package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Prettyletto/service-notifier/cmd/internal/model"
	"github.com/Prettyletto/service-notifier/cmd/internal/service"
)

type CompanyHandler struct {
	service service.CompanyService
}

func NewCompanyHandler(service service.CompanyService) *CompanyHandler {
	return &CompanyHandler{service: service}
}

func (h *CompanyHandler) CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var company model.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, fmt.Sprintf("error in the payload: %v", err), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCompany(&company); err != nil {
		http.Error(w, fmt.Sprintf("failed to create company: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CompanyHandler) GetAllCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	companys, err := h.service.RetrieveAllCompanies()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retireve companys: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(companys); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CompanyHandler) GetCompanyByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	company, err := h.service.RetrieveCompanyById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve companye :%v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *CompanyHandler) UpdateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var newCompany model.Company
	if err := json.NewDecoder(r.Body).Decode(&newCompany); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode payload: %v", err), http.StatusBadRequest)
	}

	updated, err := h.service.UpdateCompany(id, &newCompany)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update company: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	updated.ID = id

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode company: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *CompanyHandler) DeleteCollectionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.service.DeleteCompany(id); err != nil {
		http.Error(w, fmt.Sprintf("failed to delete collection: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}
