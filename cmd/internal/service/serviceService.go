package service

import (
	"github.com/Prettyletto/service-notifier/cmd/internal/model"
	"github.com/Prettyletto/service-notifier/cmd/internal/repository"
)

type ServiceService interface {
	CreateService(service *model.Service) error
	RetrieveAllServices() ([]model.Service, error)
	RetrieveServiceById(id string) (*model.Service, error)
	UpdateService(id string, service *model.Service) (*model.Service, error)
	DeleteService(id string) error
}

type serviceService struct {
	repo repository.ServiceRepository
}

func NewServiceService(repo repository.ServiceRepository) ServiceService {
	return &serviceService{repo: repo}
}

func (s *serviceService) CreateService(service *model.Service) error {

	if err := service.Validate(); err != nil {
		return err
	}

	service.GenerateID()

	return s.repo.SaveService(service)
}

func (s *serviceService) RetrieveAllServices() ([]model.Service, error) {
	return s.repo.FindAllServices()
}

func (s *serviceService) RetrieveServiceById(id string) (*model.Service, error) {

	return s.repo.FindServiceById(id)
}

func (s *serviceService) UpdateService(id string, service *model.Service) (*model.Service, error) {
	if err := service.Validate(); err != nil {
		return nil, err
	}

	err := s.repo.UpdateService(id, service)
	if err != nil {
		return nil, err
	}
	return service, err
}

func (s *serviceService) DeleteService(id string) error {
	return s.repo.DeleteService(id)
}
